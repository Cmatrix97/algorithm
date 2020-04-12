package josephus

import (
	"testing"
)

var (
	N = 10000
	M = 333
)

func BenchmarkRing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosephusRing(N, M)
	}
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosephusQueue(N, M)
	}
}

func BenchmarkForm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosephusForm1(N, M)
	}
}

func BenchmarkForm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosephusForm2(N, M)
	}
}
