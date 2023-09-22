package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Create a new Chi router.
	r := chi.NewRouter()

	// Use middleware for logging, request ID, and recovery from panics.
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// Define a route and handle function.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Start the web server on port 8080.
	http.ListenAndServe(":8080", r)
}
