package articles

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterArticlesPageServiceServer(server, &ArticlesServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestArticlesServer_GetArticles(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewArticlesPageServiceClient(conn)

	t.Run("valid results on page 1", func(t *testing.T) {
		page := int32(1)
		request := &pb.ArticlesPageRequest{
			QueryString: "Quantum",
			Page:        &page,
		}

		response, err := client.GetArticles(ctx, request)

		if err != nil {
			t.Log("server_Test.TestArticlesServer_GetArticles: error when getting articles, ", err)
		}

		// Check response has 10 results
		actualNumberOfResults := len(response.GetResults())
		if actualNumberOfResults != 10 {
			t.Error("error: expected", 10, "results received", actualNumberOfResults)
		}
	})
}
