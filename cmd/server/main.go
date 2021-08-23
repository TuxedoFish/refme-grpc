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
	"encoding/json"
)

type Config struct {
  PORT string
  SPRINGER_API_KEY string
}

func main() {
	godotenv.Load()
	configJson := os.Getenv("CONFIG")
	var config Config 
	json.Unmarshal([]byte(configJson), &config)

	PORT := config.PORT
	address := fmt.Sprintf(":%[1]s", PORT)

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
