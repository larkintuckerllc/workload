package healthv1

import (
	"context"
	"log"

	pb "github.com/larkintuckerllc/workload/pkg/healthv1"
)

type HealthServer struct {
	pb.UnimplementedHealthServer
}

func (s *HealthServer) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("Received Check")
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}
