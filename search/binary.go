// LowerBound and UpperBound returns len(a) if it failed.
// Other function returns the index of the target and whether it succeeded or not by (int, bool).
package search

// BinarySearch returns the index of target in [left, right].
func BinarySearch(a []int, left, right, target int) (int, bool) {
	for left <= right {
		mid := left + (right-left)>>1
		if target > a[mid] {
			left = mid + 1
		} else if target < a[mid] {
			right = mid - 1
		} else {
			return mid, true
		}
	}
	return -1, false
}

// LowerBound returns the index of the first element greater than or equal to target in [left, right).
func LowerBound(a []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)>>1
		if a[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// LowerBound returns the index of the first element greater than target in [left, right).
func UpperBound(a []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)>>1
		if a[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// LastLE returns the index of the last element less than or equal to the target.
func LastLE(a []int, target int) (int, bool) {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target >= a[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == 0 {
		return -1, false
	}
	return left - 1, true
}

// FirstGE returns the index of the first element greater than or equal to the target.
func FirstGE(a []int, target int) (int, bool) {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target <= a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right == len(a)-1 {
		return -1, false
	}
	return right + 1, true
}

// LastLT returns the index of the last element less than the target.
func LastLT(a []int, target int) (int, bool) {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if target > a[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == 0 {
		return -1, false
	}
	return left - 1, true
}

// FirstGT reutrns the index of the first element greater than the target.
func FirstGT(a []int, target int) (int, bool) {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right == len(a)-1 {
		return -1, false
	}
	return right + 1, true
}
