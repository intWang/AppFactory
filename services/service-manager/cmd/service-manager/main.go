package main

import (
	"log"
	"net/http"

	"appfactory/service-manager/internal/httpapi"
)

func main() {
	log.Println("service-manager listening on :8080")
	if err := http.ListenAndServe(":8080", httpapi.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
