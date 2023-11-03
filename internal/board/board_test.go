package board

import (
	"math/rand"
	"testing"
)

// V1 6.2
// V2 5.5
// V3 5.3 - Fully working
// V4 2.170 - Arrays instead of maps
// V5 2.072 - Pre-computed inverse check
func BenchmarkBoard(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	board := New(100, r)
	for !board.Iter() {
	}
}
