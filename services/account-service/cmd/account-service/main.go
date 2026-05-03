package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"appfactory/account-service/internal/httpapi"
	sharedconfig "appfactory/shared-go/config"
)

func main() {
	configPath := os.Getenv("APP_CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/local.yaml"
	}
	cfg, err := sharedconfig.LoadYAML[httpapi.Config](configPath)
	if err != nil {
		log.Fatal(err)
	}
	router, _, err := httpapi.NewPostgresRouter(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account-service listening on :%s", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		log.Fatal(err)
	}
}
