package job

import (
	pb "github.com/managef/models/rpc"
	"context"
	"fmt"
	"log"
)

type Server struct{}

// GetJob implements requests.JobServer

func (s *Server) GetJob(ctx context.Context, in *pb.JobRequest) (*pb.JobResponse, error) {
	log.Printf("Request job %+v", in)
	return &pb.JobResponse{Id: fmt.Sprintf("Hello Worker %d", in.Number)}, nil
}