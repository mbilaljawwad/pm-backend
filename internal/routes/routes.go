package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Health check route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Health check")
		w.Write([]byte("OK"))
	})
	// Health check route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Health check")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Health Check"}`))
	})

	return r
}
