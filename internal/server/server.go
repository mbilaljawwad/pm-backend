package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mbilaljawwad/pm-backend/internal/routes"
)

type Server struct {
	srv *http.Server
}

/**
 * Initialize the server with necessary configurations.
 */
func initializeServer() *http.Server {
	return &http.Server{
		Addr:           ":8081",
		Handler:        routes.SetRoutes(),
		ReadTimeout:    10 * time.Second, // Read timeout
		WriteTimeout:   10 * time.Second, // Write timeout
		MaxHeaderBytes: 1 << 20,          // Maximum header size (1 MB)
	}
}

/**
 * Create a new server instance.
 */
func NewServer() *Server {
	return &Server{
		srv: initializeServer(),
	}
}

/**
 * Start running the server.
 */
func (server *Server) Run() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting API Server!!")
		if err := server.srv.ListenAndServe(); err != nil {
			log.Fatalf("server failed to start")
		}
	}()

	// Blocking until a termination signal is received.
	<-stop
	log.Println("Preparing to shut down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.srv.Shutdown(ctx); err != nil {
		log.Fatalf("server failed to shut down %v\n", err)
	}

	log.Println("Server has been shut down successfully!")
}
