package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

var (
	rng    *rand.Rand // a handle could generates pseudo-random numbers
	length int        // length and max value of array
)

func init() {
	// completely randomly generated for reference only.
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	length = 1e4
}

type arrSorter struct {
	sortFunc func([]int)
	arr      []int
}

func execTest(sorters ...arrSorter) {
	for _, sorter := range sorters {
		sorter.sortFunc(sorter.arr)
	}
}

func convertToString(a []int) string {
	return fmt.Sprint(a)
}

func TestSort(t *testing.T) {
	var tests = []struct {
		input func(a []int)
		want  func(a []int)
	}{
		{SelectSort, DefaultSort},
		{BubbleSort, DefaultSort},
		{InsertSort, DefaultSort},
		{ShellSort, DefaultSort},
		{MergeSort, DefaultSort},
		{MergeSortBU, DefaultSort},
		{QuickSort, DefaultSort},
		{QuickSort3way, DefaultSort},
		{HeapSort, DefaultSort},
	}
	for _, test := range tests {
		arr := rng.Perm(length)
		arrCopy := make([]int, length)
		copy(arrCopy, arr)
		execTest(arrSorter{test.input, arr}, arrSorter{test.want, arrCopy})
		if got := convertToString(arr); got != convertToString(arrCopy) {
			funcName := runtime.FuncForPC(reflect.ValueOf(test.input).Pointer()).Name()
			t.Error(funcName)
		}
	}
}

func exec(sortFunc func(arr []int)) {
	arr := rng.Perm(length)
	execTest(arrSorter{sortFunc, arr})
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

func BenchmarkMergeSortBU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(MergeSortBU)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(QuickSort)
	}
}

func BenchmarkQuickSort3way(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(QuickSort3way)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec(HeapSort)
	}
}
