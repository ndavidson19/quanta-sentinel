package monitor

import (
	"bufio"
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"sentinel/internal/config"
	"sentinel/internal/metrics"
	"sentinel/internal/logparser"
)

type Monitor struct {
	dockerClient *client.Client
	config       *config.Config
}

func New(dockerClient *client.Client, cfg *config.Config) *Monitor {
	return &Monitor{
		dockerClient: dockerClient,
		config:       cfg,
	}
}

func (m *Monitor) Start() error {
	containers, err := m.dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return err
	}

	for _, container := range containers {
		go m.monitorContainer(container.ID)
	}

	return nil
}

func (m *Monitor) monitorContainer(containerID string) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true,
	}

	logs, err := m.dockerClient.ContainerLogs(context.Background(), containerID, options)
	if err != nil {
		log.Printf("Error getting logs for container %s: %v", containerID, err)
		return
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		start := time.Now()
		line := scanner.Text()
		
		_, err := logparser.ParseLog(containerID, line)
		if err != nil {
			log.Printf("Error parsing log line from container %s: %v", containerID, err)
			continue
		}

		stream := "stdout"
		if line[0] == 2 {
			stream = "stderr"
		}
		
		metrics.LogLines.WithLabelValues(containerID, stream).Inc()
		
		duration := time.Since(start).Seconds()
		metrics.ProcessingTime.WithLabelValues(containerID).Observe(duration)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning logs for container %s: %v", containerID, err)
	}
}