package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mbilaljawwad/pm-backend/internal/routes"
)

func main() {
	srv := &http.Server{
		Addr:           ":8081",
		Handler:        routes.SetRoutes(),
		ReadTimeout:    10 * time.Second, // Read timeout
		WriteTimeout:   10 * time.Second, // Write timeout
		MaxHeaderBytes: 1 << 20,          // Maximum header size (1 MB)
	}

	log.Println("Starting API server")
	srv.ListenAndServe()
	log.Println("API Server started!")

	// go func() {
	// 	log.Println("Starting API server")
	// 	srv.ListenAndServe()
	// 	log.Println("API Server started!")

	// 	// if err := srv.ListenAndServe(); err != nil {
	// 	// 	log.Fatalf("server failed to start")
	// 	// } else {
	// 	// 	log.Println("API Server started!")
	// 	// }
	// }()
}
