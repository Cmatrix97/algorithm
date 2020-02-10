package test

import (
	"algorithm/sort"
	"testing"
)

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.DefaultSort)
	}
}
func BenchmarkSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.SelectSort)
	}
}
func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.BubbleSort)
	}
}
func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.InsertSort)
	}
}
func BenchmarkShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.ShellSort)
	}
}
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.MergeSort)
	}
}
func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.HeapSort)
	}
}
func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.QuickSort)
	}
}
