package main

import (
	"log"
	"net/http"

	go_specs_greet "github.com/xandalm/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
