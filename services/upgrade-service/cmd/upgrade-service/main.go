package main

import (
	"log"
	"net/http"

	"appfactory/upgrade-service/internal/httpapi"
)

func main() {
	log.Println("upgrade-service listening on :8082")
	if err := http.ListenAndServe(":8082", httpapi.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
