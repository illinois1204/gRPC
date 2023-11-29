package provider

import (
	echo "grpc-server/pkg/echo"
	pingPong "grpc-server/pkg/ping-pong"
	"grpc-server/pkg/stream"
	"grpc-server/services"

	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	echo.RegisterEchoServer(s, &services.EchoService{})
	echo.RegisterCalculateServer(s, &services.CalculateService{})
	stream.RegisterStreamEchoServer(s, &services.StreamService{})
	pingPong.RegisterPingPongServer(s, &services.PingPongService{})
}
