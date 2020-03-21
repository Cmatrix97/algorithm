package sort

// HeapSort
func HeapSort(a []int) {
	N := len(a) - 1
	for k := (N - 1) / 2; k >= 0; k-- {
		sink(a, k, N)
	}
	for N > 0 {
		a[0], a[N] = a[N], a[0]
		N--
		sink(a, 0, N)
	}
}

func sink(pq []int, k, N int) {
	for 2*k+1 <= N {
		j := 2*k + 1
		if j < N && pq[j+1] > pq[j] {
			j++
		}
		if pq[k] >= pq[j] {
			break
		}
		pq[k], pq[j] = pq[j], pq[k]
		k = j
	}
}
