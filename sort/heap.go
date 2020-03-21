package sort

// HeapSort: Modify subscript by copy simply.
func HeapSort(a []int) {
	aux := make([]int, len(a)+1)
	copy(aux[1:], a)
	HeapS(aux)
	copy(a, aux[1:])
}

func HeapS(a []int) {
	N := len(a) - 1
	for k := N / 2; k >= 1; k-- {
		sink(a, k, N)
	}
	for N > 1 {
		a[1], a[N] = a[N], a[1]
		N--
		sink(a, 1, N)
	}
}

func swim(pq []int, k int) {
	for k > 1 && pq[k/2] < pq[k] {
		pq[k/2], pq[k] = pq[k], pq[k/2]
		k /= 2
	}
}

func sink(pq []int, k, N int) {
	for 2*k <= N {
		j := 2 * k
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
