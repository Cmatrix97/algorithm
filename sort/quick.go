package sort

// QuickSort (2way)
func QuickSort(a []int) {
	QSort(a, 0, len(a)-1)
}

func QSort(a []int, low, high int) {
	if high <= low+15 {
		InsertSort(a[low : high+1])
		return
	}
	j := partition(a, low, high)
	QSort(a, low, j-1)
	QSort(a, j+1, high)
}

func partition(a []int, low, high int) int {
	temp := a[low]
	for low < high {
		for low < high && a[high] >= temp {
			high--
		}
		a[low] = a[high]
		for low < high && a[low] <= temp {
			low++
		}
		a[high] = a[low]
	}
	a[low] = temp
	return low
}

// QuickSort3way
func QuickSort3way(a []int) {
	Quick3way(a, 0, len(a)-1)
}

func Quick3way(a []int, low, high int) {
	if high <= low {
		return
	}
	lt, gt, i := low, high, low+1
	v := a[low]
	for i <= gt {
		if a[i] < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if a[i] > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}
	Quick3way(a, low, lt-1)
	Quick3way(a, gt+1, high)
}
