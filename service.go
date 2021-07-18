package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	HOST := "0.0.0.0"
	PORT := "50051"

	fmt.Println("Starting service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("%[1]s:%[2]s", HOST, PORT))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
