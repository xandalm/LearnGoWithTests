package httpserver

import (
	"fmt"
	"net/http"

	"github.com/xandalm/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, interactions.Greet(name))
	})
	mux.HandleFunc(cursePath, func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, interactions.Curse(name))
	})
	return mux
}
