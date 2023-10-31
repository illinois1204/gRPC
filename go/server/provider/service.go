package provider

import (
	pb "grpc-server/pkg"
	"grpc-server/services"

	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterEchoServer(s, &services.EchoService{})
	pb.RegisterCalculateServer(s, &services.CalculateService{})
}
