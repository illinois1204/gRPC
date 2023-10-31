package services

import (
	"context"
	pb "grpc-server/pkg"
)

type CalculateService struct {
	pb.UnimplementedCalculateServer
}

func (s *CalculateService) Min(ctx context.Context, input *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	var a1 int32 = input.GetArg1()
	var a2 int32 = input.GetArg2()

	var result int32
	if a1 > a2 {
		result = a2
	} else {
		result = a1
	}
	return &pb.CalculateResponse{Result: result}, nil
}

func (s *CalculateService) Sum(ctx context.Context, input *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	return &pb.CalculateResponse{Result: input.GetArg1() + input.GetArg2()}, nil
}
