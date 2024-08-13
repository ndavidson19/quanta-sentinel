package docker

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil Docker client")
	}

	// Test API version negotiation
	_, err = client.ServerVersion(context.Background())
	if err != nil {
		t.Fatalf("Failed to get server version: %v", err)
	}
}