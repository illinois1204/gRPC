package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	echo "grpc-client/pkg/echo"
	"log"
	"time"
)

func main() {
	const PORT = 8717
	connect, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		echoClient      = echo.NewEchoClient(connect)
		calculateClient = echo.NewCalculateClient(connect)
	)

	echoResponse, err := echoClient.Call(ctx, &echo.EchoRequest{Name: "Jose"})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("%s\n", echoResponse.GetMessage())

	calculateResponse, err := calculateClient.Sum(ctx, &echo.CalculateRequest{Arg1: 45, Arg2: 87})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("result of sum is %d\n", calculateResponse.GetResult())

	fmt.Println("Program finished.")
}
