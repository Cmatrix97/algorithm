package test

import (
	"algorithm/math"
	"testing"
)

var (
	N = 10000
	M = 333
)

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.JosephusQueue(N, M)
	}
}

func BenchmarkForm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.JosephusForm1(N, M)
	}
}

func BenchmarkForm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.JosephusForm2(N, M)
	}
}
