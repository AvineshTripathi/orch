package api

import (
	"context"
	pb "orch/proto-provisioner"
	"orch/provisioner/queue"
)

type Server struct {
	pb.UnimplementedProvisionerServiceServer
	queueClient queue.Client
}

func (s *Server) GetProvisionerStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Data: "provisioner is up and running"}, nil
}

func (s *Server) CreateEnvironment(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	return &pb.TaskCreateResponse{Status: ""}, nil
}
