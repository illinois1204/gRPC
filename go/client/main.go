package main

import (
	"context"
	"fmt"
	echo "grpc-client/pkg/echo"
	pingPong "grpc-client/pkg/ping-pong"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	EchoStub      echo.EchoClient
	CalculateStub echo.CalculateClient
	PingPongStub  pingPong.PingPongClient
)

func makeEcho(ctx context.Context) {
	res, err := EchoStub.Call(ctx, &echo.EchoRequest{Name: "Jose"})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("%s\n", res.GetMessage())
}

func makeCalculate(ctx context.Context) {
	res, err := CalculateStub.Sum(ctx, &echo.CalculateRequest{Arg1: 45, Arg2: 87})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("result of sum is %d\n", res.GetResult())
}

func playPingPong(ctx context.Context) {
	const limit = 20
	var i uint16 = 0

	stream, err := PingPongStub.Play(ctx)
	if err != nil {
		log.Fatalf("failed to call BidirectionalStreamingEcho: %v\n", err)
	}

	fmt.Print("Ping......")
	stream.Send(&pingPong.Empty{})
	for {
		_, err := stream.Recv()
		if err != nil {
			log.Fatalf("\n%v\n", err)
			break
		}

		fmt.Println("Pong")
		i++
		if i == limit {
			stream.CloseSend()
			break
		}

		time.Sleep(500 * time.Millisecond)
		fmt.Print("Ping......")
		time.Sleep(time.Second)
		stream.Send(&pingPong.Empty{})
	}
	fmt.Println("Game over")
}

func main() {
	const PORT = 8717
	connect, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	EchoStub = echo.NewEchoClient(connect)
	CalculateStub = echo.NewCalculateClient(connect)
	PingPongStub = pingPong.NewPingPongClient(connect)

	// makeEcho(ctx)
	// makeCalculate(ctx)
	playPingPong(ctx)
	fmt.Println("\nProgram finished.")
}
