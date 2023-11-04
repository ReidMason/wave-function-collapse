package socket

import (
	"fmt"
	"testing"
)

const iters = 100000

func TestCanConnect(t *testing.T) {
	tests := [][]Socket{
		{Water, Sand},
		{WaterT, Sand},
	}
	ConvertSocketConstraints()

	t.Parallel()
	for _, test := range tests {

		r1 := CanConnect(Water, Sand)
		r2 := canConnectHashmap(Water, Sand)
		if r1 != r2 {
			t.Fatalf("Failed comparison for %s and %s", fmt.Sprint(test[0]), fmt.Sprint(test[1]))
		}
	}
}

func BenchmarkAccessConstraintsArray(b *testing.B) {
	ConvertSocketConstraints()
	b.ResetTimer()
	for i := 0; i < iters; i++ {
		CanConnect(Water, Sand)
	}
}
