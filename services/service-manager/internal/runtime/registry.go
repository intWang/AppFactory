package runtime

import (
	"appfactory/service-manager/internal/domain"
	sharedconfig "appfactory/shared-go/config"
)

type Config struct {
	ServiceName    string          `yaml:"service_name"`
	HTTPPort       string          `yaml:"http_port"`
	Environment    string          `yaml:"environment"`
	DefaultProfile string          `yaml:"default_profile"`
	Services       []ConfigService `yaml:"services"`
}

type ConfigService struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Address string `yaml:"address"`
}

type Registry struct {
	Services []domain.ServiceStatus
}

func NewRegistry() *Registry {
	return &Registry{
		Services: []domain.ServiceStatus{
			{Name: "account-service", Address: "http://localhost:8081", Status: "registered", Profile: "local"},
			{Name: "upgrade-service", Address: "http://localhost:8082", Status: "registered", Profile: "local"},
			{Name: "service-manager", Address: "http://localhost:8080", Status: "running", Profile: "local"},
		},
	}
}

func LoadRegistryFromConfig(path string) (*Registry, error) {
	cfg, err := sharedconfig.LoadYAML[Config](path)
	if err != nil {
		return nil, err
	}

	services := make([]domain.ServiceStatus, 0, len(cfg.Services))
	for _, service := range cfg.Services {
		services = append(services, domain.ServiceStatus{
			Name:    service.Name,
			Address: service.Address,
			Status:  "registered",
			Profile: cfg.DefaultProfile,
		})
	}

	return &Registry{Services: services}, nil
}
