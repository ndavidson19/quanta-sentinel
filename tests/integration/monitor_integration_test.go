package integration

import (
	"context"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
	"sentinel/internal/config"
	"sentinel/internal/monitor"
)

func TestMonitorIntegration(t *testing.T) {
	// Skip if not in integration test environment
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Create a real Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NoError(t, err)

	// Create a test container
	ctx := context.Background()
	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"sh", "-c", "while true; do echo 'Test log'; sleep 1; done"},
	}, nil, nil, nil, "")
	assert.NoError(t, err)

	// Start the container
	err = dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	assert.NoError(t, err)

	// Ensure cleanup
	defer func() {
		dockerClient.ContainerStop(ctx, resp.ID, nil)
		dockerClient.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	}()

	// Create and start the monitor
	cfg := &config.Config{}
	m := monitor.New(dockerClient, cfg)
	err = m.Start()
	assert.NoError(t, err)

	// Wait for some logs to be processed
	time.Sleep(5 * time.Second)

	// Check metrics (this is a basic check, you might want to add more specific assertions)
	// You would typically expose the metrics and check them here
	// For this example, we're just ensuring the monitor ran without error

	// If you've implemented a way to access metrics, you could add assertions here
	// For example:
	// assert.Greater(t, getMetricValue("log_lines_total"), float64(0))
}