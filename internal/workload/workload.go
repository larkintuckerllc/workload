package workload

import (
	"log"
	"net"

	"github.com/larkintuckerllc/workload/internal/greet"
	"github.com/larkintuckerllc/workload/internal/healthv1"
	greetPb "github.com/larkintuckerllc/workload/pkg/greet"
	healthV1Pb "github.com/larkintuckerllc/workload/pkg/healthv1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func Execute() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetPb.RegisterGreeterServer(s, &greet.GreeterServer{})
	healthV1Pb.RegisterHealthServer(s, &healthv1.HealthServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
