package workload

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/larkintuckerllc/workload/internal/greet"
	"github.com/larkintuckerllc/workload/internal/healthv1"
	"github.com/larkintuckerllc/workload/internal/ping"
	greetPb "github.com/larkintuckerllc/workload/pkg/greet"
	healthV1Pb "github.com/larkintuckerllc/workload/pkg/healthv1"
	pingPb "github.com/larkintuckerllc/workload/pkg/ping"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Execute() {
	var port string
	envPort := os.Getenv("PORT")
	if envPort == "" {
		port = ":50051"
	} else {
		port = fmt.Sprintf(":%s", envPort)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetPb.RegisterGreeterServer(s, &greet.GreeterServer{})
	healthV1Pb.RegisterHealthServer(s, &healthv1.HealthServer{})
	pingPb.RegisterPingerServer(s, &ping.PingerServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
