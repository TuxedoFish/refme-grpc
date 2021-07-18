package articles

import (
	"testing"

	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

func TestUtils_SplitArray(t *testing.T) {
	actual := []*pb.Provider{
		&pb.Provider{Name: "x5gon", Weight: 0, Amount: 0},
		&pb.Provider{Name: "arXiv", Weight: 0.5, Amount: 0},
		&pb.Provider{Name: "springer", Weight: 0.5, Amount: 0},
	}
	expected := []*pb.Provider{
		&pb.Provider{Name: "x5gon", Weight: 0, Amount: 0},
		&pb.Provider{Name: "arXiv", Weight: 0.5, Amount: 5},
		&pb.Provider{Name: "springer", Weight: 0.5, Amount: 5},
	}

	t.Run("production split works", func(t *testing.T) {
		splitArray(actual, 10)

		if !checkEqual(actual, expected) {
			t.Error("error: expected", expected, "received", actual)
		}
	})
}

func checkEqual(actual []*pb.Provider, expected []*pb.Provider) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i, a := range actual {
		b := expected[i]
		if b.Amount != a.Amount ||
			b.Weight != a.Weight ||
			b.Name != a.Name {
			return false
		}
	}
	return true
}
