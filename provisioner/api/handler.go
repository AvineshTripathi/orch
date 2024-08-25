package api

import (
	"context"
	pb "orch/proto-provisioner"
	"orch/provisioner/queue"
	"orch/provisioner/task"
)

type Server struct {
	pb.UnimplementedProvisionerServiceServer
	QueueClient queue.Client
}

func (s *Server) GetProvisionerStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Data: "provisioner is up and running"}, nil
}

func (s *Server) ExecuteTask(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {

	var t *task.Task
	// finding the type of config returned
	switch x := req.Config.ConfigType.(type) {
	case *pb.EnvConfig_BasicConfig:
		t = task.NewTask(x.BasicConfig.GetUrl(), x.BasicConfig.GetData())
	}

	_, err := s.QueueClient.AddNewTask(t)
	if err != nil {
		return &pb.TaskCreateResponse{Status: ""}, err
	}
	return &pb.TaskCreateResponse{Status: "OK"}, nil
}
