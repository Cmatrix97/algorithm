package sort

import "math"

var aux []int

// MergeSort: Top-down
func MergeSort(a []int) {
	aux = make([]int, len(a))
	Msort(a, 0, len(a)-1)
}

func Msort(a []int, low, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)>>1
	Msort(a, low, mid)
	Msort(a, mid+1, high)
	merge(a, low, mid, high)
}

func merge(a []int, low, mid, high int) {
	i, j := low, mid+1
	for k := low; k <= high; k++ {
		aux[k] = a[k]
	}
	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > high {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

// MergeSortBU: Bottom-up
func MergeSortBU(a []int) {
	N := len(a)
	for sz := 1; sz < N; sz += sz {
		for low := 0; low < N-sz; low += sz + sz {
			merge(a, low, low+sz-1, int(math.Min(float64(low+sz+sz-1), float64(N-1))))
		}
	}
}
