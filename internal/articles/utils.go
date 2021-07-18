package articles

import (
	"github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

/*
   Using the D'hondt method used for splitting seats
   allocate the splits of returned elements
*/
func splitArray(providers []*articlespb.Provider, n int) {
	arr := make([]int, len(providers))
	quotients := calculateQuotients(providers, arr)

	i := 0
	for i < n {
		// Get the index of the element that is largest
		idx := getLargestElementIndex(quotients)

		// Allocate 1 to the largest
		arr[idx] = arr[idx] + 1

		// Re-calculate the quotients
		quotients = calculateQuotients(providers, arr)
		i++
	}

	for i := 0; i < len(providers); i++ {
		providers[i].Amount = int32(arr[i])
	}
}

func calculateQuotients(providers []*articlespb.Provider, arr []int) []float32 {
	result := make([]float32, len(providers))

	for i := 0; i < len(providers); i++ {
		result[i] = float32(providers[i].Weight) / float32(arr[i]+1)
	}

	return result
}

func getLargestElementIndex(nums []float32) int {
	largest := float32(-1)
	largest_idx := -1

	for i, num := range nums {
		if num > largest {
			largest = num
			largest_idx = i
		}
	}

	return largest_idx
}
