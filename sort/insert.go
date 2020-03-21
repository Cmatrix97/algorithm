package sort

// InsertSort
func InsertSort(a []int) {
	N := len(a)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
