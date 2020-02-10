package sort

func ShellSort(a []int) {
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}
}
