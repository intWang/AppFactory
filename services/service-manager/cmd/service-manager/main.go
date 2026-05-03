package main

import (
	"log"
	"net/http"

	"appfactory/service-manager/internal/httpapi"
	"appfactory/service-manager/internal/runtime"
)

func main() {
	manager, err := runtime.NewManagerFromConfig("configs/local.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("service-manager listening on :8080")
	if err := http.ListenAndServe(":8080", httpapi.NewRouterWithManager(manager)); err != nil {
		log.Fatal(err)
	}
}
