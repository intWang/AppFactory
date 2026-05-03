package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	ServiceName string `yaml:"service_name"`
	HTTPPort    string `yaml:"http_port"`
	Environment string `yaml:"environment"`
}

func LoadYAML[T any](path string) (T, error) {
	var cfg T
	content, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
