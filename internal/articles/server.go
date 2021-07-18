package articles

import (
	"context"
	"fmt"

	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

type ArticlesServer struct {
	pb.UnimplementedArticlesPageServiceServer
}

func (*ArticlesServer) GetArticles(ctx context.Context, req *pb.ArticlesPageRequest) (*pb.ArticlesPageResponse, error) {
	fmt.Printf("GetArticles called with: %v", req)
	return nil, nil
}
