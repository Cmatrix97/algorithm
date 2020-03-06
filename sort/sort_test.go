package sort

import (
	"math/rand"
	"testing"
	"time"
)

var (
	rng    *rand.Rand // a handle could generates pseudo-random numbers
	length int        // length and max value of array
)

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	length = 1e4
}

func exec(sortFunc func(arr []int)) {
	arr := rng.Perm(length)
	sortFunc(arr)
}

func BenchmarkDefaultSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(DefaultSort)
	}
}
func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(SelectSort)
	}
}
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(BubbleSort)
	}
}
func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(InsertSort)
	}
}
func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(ShellSort)
	}
}
func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(MergeSort)
	}
}
func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(HeapSort)
	}
}
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(QuickSort)
	}
}
