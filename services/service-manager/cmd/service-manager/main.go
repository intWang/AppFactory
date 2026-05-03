package main

import (
	"log"
	"net/http"
	"os"

	"appfactory/service-manager/internal/httpapi"
	"appfactory/service-manager/internal/runtime"
	sharedconfig "appfactory/shared-go/config"
)

func main() {
	configPath := os.Getenv("APP_CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/local.yaml"
	}
	manager, err := runtime.NewManagerFromConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := sharedconfig.LoadYAML[httpapi.Config](configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("service-manager listening on :8080")
	if err := http.ListenAndServe(":8080", httpapi.NewRouterWithConfig(manager, cfg)); err != nil {
		log.Fatal(err)
	}
}
