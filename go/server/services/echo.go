package services

import (
	"context"
	"fmt"
	"grpc-server/pkg/echo"
	"time"
)

type EchoService struct {
	echo.UnimplementedEchoServer
}

func (s *EchoService) Call(ctx context.Context, input *echo.EchoRequest) (*echo.EchoResponse, error) {
	var name string = input.GetName()
	return &echo.EchoResponse{Message: fmt.Sprintf("yaos, %s, this response from server. Time is %s", name, time.Now().String())}, nil
}
