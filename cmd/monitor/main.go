package main

import (
	"log"
	"net/http"

	"sentinel/internal/config"
	"sentinel/internal/docker"
	"sentinel/internal/metrics"
	"sentinel/internal/monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Docker client
	dockerClient, err := docker.NewClient()
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	// Initialize metrics
	metrics.Init()

	// Start Prometheus HTTP server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Starting metrics server on %s", cfg.MetricsAddr)
		if err := http.ListenAndServe(cfg.MetricsAddr, nil); err != nil {
			log.Fatalf("Error starting metrics server: %v", err)
		}
	}()

	// Start the monitor
	m := monitor.New(dockerClient, cfg)
	if err := m.Start(); err != nil {
		log.Fatalf("Error starting monitor: %v", err)
	}

	// Keep the main goroutine alive
	select {}
}