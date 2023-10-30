package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-server/pkg"

	"google.golang.org/grpc"
)

type EchoService struct {
	pb.UnimplementedEchoServer
}

func (s *EchoService) Call(ctx context.Context, input *pb.EchoRequest) (*pb.EchoResponse, error) {
	var name string = input.GetName()
	return &pb.EchoResponse{Message: fmt.Sprintf("yaos, %s, this response from server. Time is %s", name, time.Now().String())}, nil
}

func main() {
	const PORT uint16 = 8718
	address, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &EchoService{})

	fmt.Printf("Server run on port: %d\n", PORT)
	if err := s.Serve(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
