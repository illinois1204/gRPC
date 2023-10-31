package services

import (
	"context"
	"fmt"
	pb "grpc-server/pkg"
	"time"
)

type EchoService struct {
	pb.UnimplementedEchoServer
}

func (s *EchoService) Call(ctx context.Context, input *pb.EchoRequest) (*pb.EchoResponse, error) {
	var name string = input.GetName()
	return &pb.EchoResponse{Message: fmt.Sprintf("yaos, %s, this response from server. Time is %s", name, time.Now().String())}, nil
}
