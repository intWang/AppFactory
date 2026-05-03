package main

import (
	"log"
	"net/http"

	"appfactory/account-service/internal/httpapi"
)

func main() {
	log.Println("account-service listening on :8081")
	if err := http.ListenAndServe(":8081", httpapi.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
