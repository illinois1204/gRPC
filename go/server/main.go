package main

import (
	"fmt"
	"grpc-server/provider"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	const PORT uint16 = 8718
	address, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	provider.RegisterServices(s)

	fmt.Printf("Server run on port: %d\n", PORT)
	if err := s.Serve(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
