package main

import (
	"context"
	"fmt"
	pb "grpc-client/pkg"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connect, err := grpc.Dial(":8717", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	var (
		echoClient = pb.NewEchoClient(connect)
		calculateClient = pb.NewCalculateClient(connect)
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	echoResponse, err := echoClient.Call(ctx, &pb.EchoRequest{Name: "Jose"})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("%s\n", echoResponse.GetMessage())

	calculateResponse, err := calculateClient.Sum(ctx, &pb.CalculateRequest{Arg1: 45, Arg2: 87})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Printf("result of sum is %d\n", calculateResponse.GetResult())

	fmt.Println("Program finished.")
}
