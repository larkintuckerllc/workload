package workload

import (
	"log"
	"net"

	"github.com/larkintuckerllc/workload/internal/greet"
	greetPb "github.com/larkintuckerllc/workload/pkg/greet"
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
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
