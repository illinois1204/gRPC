package services

import (
	"fmt"
	"grpc-server/pkg/stream"
	"io"
	"time"
)

type StreamService struct {
	stream.UnimplementedStreamEchoServer
}

func (s *StreamService) RunServerStream(in *stream.RequestPayload, call stream.StreamEcho_RunServerStreamServer) error {
	fmt.Printf("Incoming request -> code: %d, msg: %s\n", in.GetCode(), in.GetMsg())
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i + 1);
        fmt.Println("processing...");
		time.Sleep(1500 * time.Millisecond)
		call.Send(&stream.ResponsePayload{Code: 1, Status: "CONTINUE"})
	}
	fmt.Println("Finished processing");
	call.Send(&stream.ResponsePayload{Code: 0, Status: "END"})
	return nil
}

func (s *StreamService) RunClientStream(call stream.StreamEcho_RunClientStreamServer) error {
	var arr []*stream.ResponsePayload
	for {
		value, err := call.Recv()
		if err == io.EOF {
			return call.SendAndClose(&stream.ArrResponsePayload{Info: arr})
		}
		if err != nil {
			return err
		}

		fmt.Printf("Incoming request -> %v\n", value)
		var request = stream.ResponsePayload{Code: value.GetCode(), Status: value.GetMsg()}
		arr = append(arr, &request)
	}
}

func (s *StreamService) RunDuplexStream(call stream.StreamEcho_RunDuplexStreamServer) error {
	var iterator uint8 = 0

	for {
		value, err := call.Recv()
		if err == io.EOF {
			call.Send(&stream.ResponsePayload{Code: 0, Status: "END"})
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf("Incoming request -> %v\n", value)
		if value.GetCode() > 0 {
			iterator++
			call.Send(&stream.ResponsePayload{Code: 1, Status: "CONTINUE"})
		}
		if iterator == 10 {
			call.Send(&stream.ResponsePayload{Code: 2, Status: "FINISH"})
			return nil
		}
	}
}
