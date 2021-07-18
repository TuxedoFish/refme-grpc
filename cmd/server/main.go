package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetArticles(ctx context.Context, req *articlespb.ArticlesPageRequest) (*articlespb.ArticlesPageResponse, error) {
	fmt.Printf("GetArticles called with: %v", req)
	return nil, nil
}

func main() {
	HOST := "0.0.0.0"
	PORT := "50051"
	address := fmt.Sprintf("%[1]s:%[2]s", HOST, PORT)

	fmt.Printf("Starting service (%v)...\n", address)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	articlespb.RegisterArticlesPageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
