package api

import (
	"context"
	pb "github.com/AvineshTripathi/orch/proto-provisioner"
	"github.com/AvineshTripathi/orch/provisioner/queue"
	"github.com/AvineshTripathi/orch/provisioner/task"
)

type Server struct {
	pb.UnimplementedProvisionerServiceServer
	QueueClient queue.Client
}

func (s *Server) GetProvisionerStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Data: "provisioner is up and running"}, nil
}

func (s *Server) ExecuteTask(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	
	t := task.NewTask(req.Name, "", req.Data)

	_, err := s.QueueClient.AddNewTask(t)
	if err != nil {
		return &pb.TaskCreateResponse{Status: ""}, err
	}
	return &pb.TaskCreateResponse{Status: "OK"}, nil
}
