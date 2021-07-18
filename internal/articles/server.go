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
	fmt.Printf("GetArticles called with: %v \n", req)

	// Setup the providers
	providers := []*pb.Provider{
		&pb.Provider{Name: "x5gon", Weight: 0, Amount: 0},
		&pb.Provider{Name: "arXiv", Weight: 0.5, Amount: 0},
		&pb.Provider{Name: "springer", Weight: 0.5, Amount: 0},
	}
	// Use D'Hondt method to split
	splitArray(providers, 10)

	res := pb.ArticlesPageResponse{
		Meta: &pb.Meta{
			Providers: providers,
			Query:     req.QueryString,
			Page:      *req.Page,
			Results:   1,
		},
		Results: []*pb.Result{
			&pb.Result{
				Id:            "Test",
				Author:        "Test",
				Title:         "Test",
				PublishedDate: "Test",
				Publisher:     "Test",
				Description:   "Test",
				Url:           "Test",
			},
		},
	}

	return &res, nil
}
