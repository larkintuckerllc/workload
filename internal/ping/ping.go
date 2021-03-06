package ping

import (
	"context"
	"log"

	pb "github.com/larkintuckerllc/workload/pkg/ping"
)

type PingerServer struct {
	pb.UnimplementedPingerServer
}

func (s *PingerServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received Ping")
	return &pb.PingReply{Message: "Pong"}, nil
}
