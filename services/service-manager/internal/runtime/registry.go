package runtime

import "appfactory/service-manager/internal/domain"

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
