package runtime

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

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
	Name          string `yaml:"name"`
	Command       string `yaml:"command"`
	WorkDir       string `yaml:"workdir"`
	Address       string `yaml:"address"`
	ContainerName string `yaml:"container_name"`
}

type Registry struct {
	Services []domain.ServiceStatus
}

type ManagedService struct {
	Name          string    `json:"name"`
	Command       string    `json:"command"`
	WorkDir       string    `json:"workdir"`
	Address       string    `json:"address"`
	ContainerName string    `json:"container_name,omitempty"`
	Status        string    `json:"status"`
	Profile       string    `json:"profile"`
	PID           int       `json:"pid"`
	StartedAt     time.Time `json:"started_at"`
	LastError     string    `json:"last_error,omitempty"`
}

type managedProcess struct {
	cmd           *exec.Cmd
	done          chan error
	stopRequested bool
}

type Manager struct {
	profile       string
	mu            sync.Mutex
	services      map[string]*ManagedService
	processes     map[string]*managedProcess
	commandRunner func(name string, args ...string) error
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

func NewManager(configs []ConfigService, profile string) *Manager {
	services := make(map[string]*ManagedService, len(configs))
	for _, service := range configs {
		copied := service
		services[service.Name] = &ManagedService{
			Name:          copied.Name,
			Command:       copied.Command,
			WorkDir:       copied.WorkDir,
			Address:       copied.Address,
			ContainerName: copied.ContainerName,
			Status:        "registered",
			Profile:       profile,
		}
	}
	return &Manager{
		profile:       profile,
		services:      services,
		processes:     map[string]*managedProcess{},
		commandRunner: runCommand,
	}
}

func NewManagerFromConfig(path string) (*Manager, error) {
	cfg, err := sharedconfig.LoadYAML[Config](path)
	if err != nil {
		return nil, err
	}
	return NewManager(cfg.Services, cfg.DefaultProfile), nil
}

func (m *Manager) List() []ManagedService {
	m.mu.Lock()
	defer m.mu.Unlock()
	results := make([]ManagedService, 0, len(m.services))
	for _, service := range m.services {
		results = append(results, *service)
	}
	return results
}

func (m *Manager) Start(name string) (ManagedService, error) {
	m.mu.Lock()
	service, ok := m.services[name]
	if !ok {
		m.mu.Unlock()
		return ManagedService{}, fmt.Errorf("unknown service: %s", name)
	}
	if process, exists := m.processes[name]; exists && process.cmd.Process != nil {
		result := *service
		m.mu.Unlock()
		return result, nil
	}
	if service.ContainerName != "" {
		containerName := service.ContainerName
		m.mu.Unlock()
		if err := m.commandRunner("docker", "start", containerName); err != nil {
			m.mu.Lock()
			service.LastError = err.Error()
			service.Status = "error"
			result := *service
			m.mu.Unlock()
			return result, err
		}
		m.mu.Lock()
		service.Status = "running"
		service.LastError = ""
		service.StartedAt = time.Now().UTC()
		result := *service
		m.mu.Unlock()
		return result, nil
	}
	commandParts := strings.Fields(service.Command)
	if len(commandParts) == 0 {
		m.mu.Unlock()
		return ManagedService{}, fmt.Errorf("service %s has empty command", name)
	}
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Dir = service.WorkDir
	if err := cmd.Start(); err != nil {
		service.LastError = err.Error()
		service.Status = "error"
		result := *service
		m.mu.Unlock()
		return result, err
	}
	service.Status = "running"
	service.LastError = ""
	service.PID = cmd.Process.Pid
	service.StartedAt = time.Now().UTC()
	process := &managedProcess{
		cmd:  cmd,
		done: make(chan error, 1),
	}
	m.processes[name] = process
	result := *service
	m.mu.Unlock()

	go m.waitForExit(name, process)

	return result, nil
}

func (m *Manager) Stop(name string) (ManagedService, error) {
	m.mu.Lock()
	service, ok := m.services[name]
	if !ok {
		m.mu.Unlock()
		return ManagedService{}, fmt.Errorf("unknown service: %s", name)
	}
	process, exists := m.processes[name]
	if service.ContainerName != "" {
		containerName := service.ContainerName
		m.mu.Unlock()
		if err := m.commandRunner("docker", "stop", containerName); err != nil {
			m.mu.Lock()
			service.LastError = err.Error()
			service.Status = "error"
			result := *service
			m.mu.Unlock()
			return result, err
		}
		m.mu.Lock()
		service.Status = "stopped"
		service.LastError = ""
		result := *service
		m.mu.Unlock()
		return result, nil
	}
	if !exists || process.cmd.Process == nil {
		service.Status = "stopped"
		result := *service
		m.mu.Unlock()
		return result, nil
	}
	cmd := process.cmd
	proc := cmd.Process
	process.stopRequested = true
	m.mu.Unlock()

	_ = proc.Signal(syscall.SIGINT)

	select {
	case <-process.done:
	case <-time.After(2 * time.Second):
		_ = proc.Kill()
		<-process.done
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	result := *service
	return result, nil
}

func (m *Manager) Restart(name string) (ManagedService, error) {
	m.mu.Lock()
	service, ok := m.services[name]
	if !ok {
		m.mu.Unlock()
		return ManagedService{}, fmt.Errorf("unknown service: %s", name)
	}
	if service.ContainerName != "" {
		containerName := service.ContainerName
		m.mu.Unlock()
		if err := m.commandRunner("docker", "restart", containerName); err != nil {
			m.mu.Lock()
			service.LastError = err.Error()
			service.Status = "error"
			result := *service
			m.mu.Unlock()
			return result, err
		}
		m.mu.Lock()
		service.Status = "running"
		service.LastError = ""
		service.StartedAt = time.Now().UTC()
		result := *service
		m.mu.Unlock()
		return result, nil
	}
	m.mu.Unlock()
	if _, err := m.Stop(name); err != nil {
		return ManagedService{}, err
	}
	return m.Start(name)
}

func (m *Manager) SwitchProfile(profile string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.profile = profile
	for _, service := range m.services {
		service.Profile = profile
	}
}

func (m *Manager) waitForExit(name string, process *managedProcess) {
	err := process.cmd.Wait()

	m.mu.Lock()
	service, ok := m.services[name]
	if !ok {
		m.mu.Unlock()
		process.done <- err
		close(process.done)
		return
	}
	current, exists := m.processes[name]
	if !exists || current != process {
		m.mu.Unlock()
		process.done <- err
		close(process.done)
		return
	}
	delete(m.processes, name)
	service.PID = 0
	if process.stopRequested {
		service.Status = "stopped"
		service.LastError = ""
	} else if err != nil && !isInterrupted(err) {
		service.Status = "error"
		service.LastError = err.Error()
	} else {
		service.Status = "stopped"
		service.LastError = ""
	}
	m.mu.Unlock()

	process.done <- err
	close(process.done)
}

func isInterrupted(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "signal: interrupt"
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}
