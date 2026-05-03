package runtime

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadRegistryFromConfigParsesServices(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "local.yaml")
	content := []byte("" +
		"service_name: service-manager\n" +
		"http_port: \"8080\"\n" +
		"environment: local\n" +
		"default_profile: local\n" +
		"services:\n" +
		"  - name: account-service\n" +
		"    command: go run ./cmd/account-service\n" +
		"    address: http://localhost:8081\n")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	registry, err := LoadRegistryFromConfig(path)
	if err != nil {
		t.Fatalf("expected registry to load, got error: %v", err)
	}

	if len(registry.Services) != 1 {
		t.Fatalf("expected one service entry, got %d", len(registry.Services))
	}
	if registry.Services[0].Name != "account-service" || registry.Services[0].Address != "http://localhost:8081" {
		t.Fatalf("unexpected service entry: %+v", registry.Services[0])
	}
}

func TestManagerCanStartAndStopLocalProcess(t *testing.T) {
	manager := NewManager([]ConfigService{
		{
			Name:    "sleepy",
			Command: "sleep 30",
			WorkDir: ".",
			Address: "http://localhost:65535",
		},
	}, "local")

	started, err := manager.Start("sleepy")
	if err != nil {
		t.Fatalf("expected process to start, got error: %v", err)
	}
	if started.PID == 0 || started.Status != "running" {
		t.Fatalf("unexpected started service state: %+v", started)
	}

	stopped, err := manager.Stop("sleepy")
	if err != nil {
		t.Fatalf("expected process to stop, got error: %v", err)
	}
	if stopped.Status != "stopped" {
		t.Fatalf("expected stopped status, got %+v", stopped)
	}
}

func TestManagerRestartReturnsFreshProcessMetadata(t *testing.T) {
	manager := NewManager([]ConfigService{
		{
			Name:    "sleepy",
			Command: "sleep 30",
			WorkDir: ".",
			Address: "http://localhost:65535",
		},
	}, "local")

	first, err := manager.Start("sleepy")
	if err != nil {
		t.Fatalf("expected process to start, got error: %v", err)
	}
	time.Sleep(50 * time.Millisecond)

	restarted, err := manager.Restart("sleepy")
	if err != nil {
		t.Fatalf("expected process to restart, got error: %v", err)
	}
	if restarted.Status != "running" {
		t.Fatalf("expected running status after restart, got %+v", restarted)
	}
	if restarted.StartedAt == first.StartedAt {
		t.Fatalf("expected restart to refresh started timestamp")
	}

	_, _ = manager.Stop("sleepy")
}
