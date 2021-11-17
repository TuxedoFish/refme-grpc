package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/TuxedoFish/refme-grpc/internal/articles"
	"github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	address := fmt.Sprintf("0.0.0.0:%[1]s", PORT)

	fmt.Printf("Starting service (%v)...\n", address)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	articlespb.RegisterArticlesPageServiceServer(s, &articles.ArticlesServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
