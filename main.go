package main

import (
	"context"
	"log"
	"net/http"
	"orch/config"
	"orch/handlers"
	"orch/middleware"
	"orch/provisioner-client"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	config.Load()

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	// Define routes
	r.HandleFunc("/", handlers.ApiServerStatusHandler).Methods(http.MethodGet)
	r.HandleFunc("/provisioner", handlers.ProvisionerStatusHandler).Methods(http.MethodGet)

	// Set up server
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,
	}

	// Run server in a goroutine
	go func() {
		log.Printf("Server is running on port %s", config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	provisioner.InitializeClient()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
