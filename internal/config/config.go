package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Network NetworkConfig `yaml:"network"`
}

type NetworkConfig struct {
	StartIP string `yaml:"start_ip"`
	EndIP   string `yaml:"end_ip"`
	Subnet  string `yaml:"subnet"`
}

func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}

func LoadDefault() (*Config, error) {
	return Load("config.yaml")
}
