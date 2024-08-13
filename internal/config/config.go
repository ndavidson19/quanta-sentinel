package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MetricsAddr string
	DockerHost  string
}

func Load() (*Config, error) {
	viper.SetDefault("METRICS_ADDR", ":8080")
	viper.SetDefault("DOCKER_HOST", "unix:///var/run/docker.sock")

	viper.AutomaticEnv()

	cfg := &Config{
		MetricsAddr: viper.GetString("METRICS_ADDR"),
		DockerHost:  viper.GetString("DOCKER_HOST"),
	}

	return cfg, nil
}