package articles

import (
	"context"
	"fmt"

	"github.com/TuxedoFish/refme-grpc/internal/articles/apis"
	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

type ArticlesServer struct {
	pb.UnimplementedArticlesPageServiceServer
}

func (*ArticlesServer) GetArticles(ctx context.Context, req *pb.ArticlesPageRequest) (*pb.ArticlesPageResponse, error) {
	fmt.Printf("GetArticles called with: %v \n", req)
	query := req.QueryString
	page := req.Page
	results := make([]*pb.Result, 0)

	// Setup the providers
	providers := []*pb.Provider{
		&pb.Provider{Name: "x5gon", Weight: 0, Amount: 0},
		&pb.Provider{Name: "arXiv", Weight: 0.5, Amount: 0},
		&pb.Provider{Name: "springer", Weight: 0.5, Amount: 0},
	}
	// Use D'Hondt method to split
	splitArray(providers, 10)
	// Get results from the three providers
	if providers[0].Amount != 0 {
		// X5GON Request (Deprecated)
	}
	if providers[1].Amount != 0 {
		// ArXiv Request
		results = append(results, apis.GetArXivArticles(query, int(providers[1].Amount), int(*page))...)
	}
	if providers[2].Amount != 0 {
		// Springer Request
	}
	numberOfResults := int32(len(results))

	res := pb.ArticlesPageResponse{
		Meta: &pb.Meta{
			Providers: providers,
			Query:     query,
			Page:      *page,
			Results:   numberOfResults,
		},
		Results: results,
	}

	return &res, nil
}
