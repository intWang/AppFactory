package main

import (
	"context"
	"log"
	"net/http"

	sharedconfig "appfactory/shared-go/config"
	"appfactory/upgrade-service/internal/httpapi"
)

func main() {
	cfg, err := sharedconfig.LoadYAML[httpapi.Config]("configs/local.yaml")
	if err != nil {
		log.Fatal(err)
	}
	router, _, err := httpapi.NewPostgresRouter(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("upgrade-service listening on :%s", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		log.Fatal(err)
	}
}
