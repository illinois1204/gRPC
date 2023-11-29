package services

import (
	"context"
	"grpc-server/pkg/echo"
)

type CalculateService struct {
	echo.UnimplementedCalculateServer
}

func (s *CalculateService) Min(ctx context.Context, input *echo.CalculateRequest) (*echo.CalculateResponse, error) {
	var a1 int32 = input.GetArg1()
	var a2 int32 = input.GetArg2()

	var result int32
	if a1 > a2 {
		result = a2
	} else {
		result = a1
	}
	return &echo.CalculateResponse{Result: result}, nil
}

func (s *CalculateService) Sum(ctx context.Context, input *echo.CalculateRequest) (*echo.CalculateResponse, error) {
	return &echo.CalculateResponse{Result: input.GetArg1() + input.GetArg2()}, nil
}
