package sort

func SelectSort(a []int) {
	N := len(a)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
