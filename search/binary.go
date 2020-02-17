/*
Each function returns the index of the target and whether it succeeded or not by (int, bool). 
*/
package search

/**
find target.
*/
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

/**
find the last element less than or equal to the target.
*/
func FindLastLE(a []int, left, right, target int) (int, bool) {
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

/**
find the first element greater than or equal to the target.
*/
func FindFirstGE(a []int, left, right, target int) (int, bool) {
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

/**
find the last element less than the target.
*/
func FindLastLT(a []int, left, right, target int) (int, bool) {
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

/**
find the first element greater than the target.
*/
func FindFirstGT(a []int, left, right, target int) (int, bool) {
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
