package api

import (
	"context"
	pb "orch/proto-provisioner"
)

// server implements provisioner.ProvisionerServiceServer
type Server struct {
	pb.UnimplementedProvisionerServiceServer
}

// GetProvisionerStatus implements provisioner.ProvisionerServiceServer
func (s *Server) GetProvisionerStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Data: "provisioner is up and running"}, nil
}

func (s *Server) CreateEnvironment(ctx context.Context, req *pb.EnvCreateRequest) (*pb.EnvCreateResponse, error) {
	return &pb.EnvCreateResponse{Status: "Environment creation started"}, nil
}
