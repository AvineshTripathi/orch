package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "orch/proto-provisioner"
	"orch/provisioner/api"

	"google.golang.org/grpc"
)

func main() {
	// Set up listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterProvisionerServiceServer(grpcServer, &api.Server{})

	// Channel for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Run gRPC server in a goroutine
	go func() {
		log.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	<-quit
	log.Println("Shutting down gRPC server...")

	// Graceful stop
	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")

	log.Println("Server exited cleanly")
}
