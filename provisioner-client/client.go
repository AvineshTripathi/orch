package provisioner

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "orch/proto-provisioner"
)

// Declare a global variable for the gRPC client
var client pb.ProvisionerServiceClient

// InitializeClient initializes the gRPC client
func InitializeClient() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewProvisionerServiceClient(conn)
}

// GetProvisionerStatus sends a request to the gRPC server and returns the response
func GetProvisionerStatus() (string, error) {
	resp, err := client.GetProvisionerStatus(context.Background(), &pb.StatusRequest{})
	if err != nil {
		return "", err
	}
	return resp.GetData(), nil
}
