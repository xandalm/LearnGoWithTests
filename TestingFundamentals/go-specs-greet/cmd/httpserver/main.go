package main

import (
	"log"
	"net/http"

	"github.com/xandalm/go-specs-greet/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
