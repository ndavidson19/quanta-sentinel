package monitor

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"sentinel/internal/config"
)

type MockDockerClient struct {
	mock.Mock
}

func (m *MockDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]types.Container), args.Error(1)
}

func (m *MockDockerClient) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	args := m.Called(ctx, container, options)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func TestMonitorStart(t *testing.T) {
	mockClient := new(MockDockerClient)
	cfg := &config.Config{}

	containers := []types.Container{
		{ID: "container1"},
		{ID: "container2"},
	}

	mockClient.On("ContainerList", mock.Anything, mock.Anything).Return(containers, nil)
	mockClient.On("ContainerLogs", mock.Anything, mock.Anything, mock.Anything).Return(io.NopCloser(strings.NewReader("")), nil)

	monitor := New(mockClient, cfg)

	err := monitor.Start()

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}