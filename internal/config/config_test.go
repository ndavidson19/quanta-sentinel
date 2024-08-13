package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("METRICS_ADDR", ":9090")
	os.Setenv("DOCKER_HOST", "tcp://localhost:2375")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	// Check if the configuration values are correctly loaded
	if cfg.MetricsAddr != ":9090" {
		t.Errorf("Expected MetricsAddr to be ':9090', got '%s'", cfg.MetricsAddr)
	}

	if cfg.DockerHost != "tcp://localhost:2375" {
		t.Errorf("Expected DockerHost to be 'tcp://localhost:2375', got '%s'", cfg.DockerHost)
	}

	// Test default values
	os.Unsetenv("METRICS_ADDR")
	os.Unsetenv("DOCKER_HOST")

	cfg, err = Load()
	if err != nil {
		t.Fatalf("Failed to load configuration with default values: %v", err)
	}

	if cfg.MetricsAddr != ":8080" {
		t.Errorf("Expected default MetricsAddr to be ':8080', got '%s'", cfg.MetricsAddr)
	}

	if cfg.DockerHost != "unix:///var/run/docker.sock" {
		t.Errorf("Expected default DockerHost to be 'unix:///var/run/docker.sock', got '%s'", cfg.DockerHost)
	}
}