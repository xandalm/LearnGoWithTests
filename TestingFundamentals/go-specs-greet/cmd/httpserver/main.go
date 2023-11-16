package main

import (
	"log"
	"net/http"

	"github.com/xandalm/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := httpserver.NewHandler()
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
