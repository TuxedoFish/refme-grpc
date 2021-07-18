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
)

func main() {
	godotenv.Load()
	HOST := os.Getenv("SVC_HOST")
	PORT := os.Getenv("SVC_PORT")
	address := fmt.Sprintf("%[1]s:%[2]s", HOST, PORT)

	fmt.Printf("Starting service (%v)...\n", address)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	articlespb.RegisterArticlesPageServiceServer(s, &articles.ArticlesServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
