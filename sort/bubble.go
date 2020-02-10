package sort

func BubbleSort(a []int) {
	N := len(a)
	for i := N - 1; i >= 1; i-- {
		flag := false
		for j := 1; j <= i; j++ {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
