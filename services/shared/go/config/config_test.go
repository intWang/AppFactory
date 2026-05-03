package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadYAMLDecodesTypedConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	content := []byte("service_name: demo\nhttp_port: \"8080\"\nenvironment: local\n")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	cfg, err := LoadYAML[AppConfig](path)
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.ServiceName != "demo" || cfg.HTTPPort != "8080" || cfg.Environment != "local" {
		t.Fatalf("unexpected config loaded: %+v", cfg)
	}
}
