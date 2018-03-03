package job

import (
	pb "github.com/managef/models/rpc"
	"context"
)

type server struct{}

// GetJob implements job.JobServer

func (s *server) GetJob(ctx context.Context, in *pb.JobRequest) (*pb.JobResponse, error) {
	return &pb.JobResponse{Id: "Hello Worker "}, nil
}