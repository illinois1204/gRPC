package services

import (
	"fmt"
	pingPong "grpc-server/pkg/ping-pong"
	"io"
)

type PingPongService struct {
	pingPong.UnimplementedPingPongServer
}

func (s *PingPongService) Play(call pingPong.PingPong_PlayServer) error {
	for {
		_, err := call.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		
		fmt.Println("accept ping")
		call.Send(&pingPong.Empty{})
	}
}
