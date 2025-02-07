package main

import (
	"log"
	"net"
	pb "github.com/AvineshTripathi/orch/proto-provisioner"
	config "github.com/AvineshTripathi/orch/config"
	"github.com/AvineshTripathi/orch/provisioner/api"
	"github.com/AvineshTripathi/orch/provisioner/queue"
	"github.com/AvineshTripathi/orch/provisioner/task"
	"github.com/AvineshTripathi/orch/provisioner/workers"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

func main() {

	var wg sync.WaitGroup
	taskChan := make(chan task.Task, 10)
	errChan := make(chan task.Task, 5)
	config.Load()

	taskExecutor := task.NewTaskExecutor()
	wk := workers.NewWorker(1, taskChan, errChan, taskExecutor, &wg)

	err := taskExecutor.LoadPlugin(config.PluginsPath)
	if err != nil {
		log.Fatalf("Cannot load plugin: %v", err)
	}

	wk.StartWorker()

	queueClient := queue.NewConnection()
	queueClient.ConfigureTaskChannel(taskChan, errChan)

	queueClient.ProcessTasksContinuously()
	queueClient.ProcessErrorConitnuously()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	apiServer := &api.Server{
		QueueClient: *queueClient,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProvisionerServiceServer(grpcServer, apiServer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down gRPC server...")

	queueClient.StopClient()
	wk.StopWorker()

	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")

	log.Println("Server exited cleanly")
}
