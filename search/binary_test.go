package search

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	maxLen = 100
	maxVal = 200
)

func randomArray(rng *rand.Rand) (arr []int, target int, res []int, suc bool) {
	n := rng.Intn(maxLen)
	arr = make([]int, n)
	for i := range arr {
		arr[i] = rng.Intn(maxVal)
	}
	sort.Ints(arr)
	target = rand.Intn(maxLen)
	res, suc = linearSearch(arr, target)
	return
}

func linearSearch(arr []int, target int) ([]int, bool) {
	var res []int
	for i, v := range arr {
		if v == target {
			res = append(res, i)
		} else if v > target {
			break
		}
	}
	if len(res) == 0 {
		return nil, false
	}
	return res, true
}

func TestBinarySearch(t *testing.T) {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10000; i++ {
		arr, target, wantIdxs, wantSuc := randomArray(rng)
		idx, suc := BinarySearch(arr, target)
		flag := false
		for _, v := range wantIdxs {
			if v == idx {
				flag = true
				break
			}
		}
		if suc != wantSuc || !flag {
			if len(wantIdxs) == 0 && idx == -1 {
				continue
			}
			t.Errorf("BinarySearch(%v, %d): got(%d, %t), want(%v, %t)", arr, target, idx, suc, wantIdxs, wantSuc)
		}
	}
}

func TestFindLastLE(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		idx    int
		suc    bool
	}{
		{[]int{63, 161, 171, 376, 472}, 240, 2, true},
		{[]int{12, 34, 118, 168, 214, 243, 314, 403, 485, 493, 493}, 214, 4, true},
		{[]int{88, 140, 144, 160, 174, 179, 242, 248, 369, 457, 467}, 242, 6, true},
		{[]int{33, 162, 192, 245, 276, 289, 430, 436}, 1306, 7, true},
		{[]int{47, 128, 153, 170, 305, 320, 359, 424, 458, 689, 749, 780, 830, 877, 896, 901, 989}, 458, 8, true},
	}
	for _, test := range tests {
		idx, suc := FindLastLE(test.arr, test.target)
		if idx != test.idx || suc != test.suc {
			t.Errorf("FindLastLE(%v, %d): got(%d,%t), except(%d,%t)", test.arr, test.target, idx, suc, test.idx, test.suc)
		}
	}
}

func TestFindFirstGE(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		idx    int
		suc    bool
	}{
		{[]int{35, 93, 163, 276, 333, 347, 415, 476}, 207, 3, true},
		{[]int{50, 66, 74, 135, 147, 158, 163, 196, 199, 223, 291, 293, 307, 349, 363, 444, 471, 493}, 163, 6, true},
		{[]int{39, 53, 76, 115, 155, 182, 182, 192, 196, 261, 276, 386, 410, 428, 438, 470, 479}, 192, 7, true},
		{[]int{31, 130, 135, 139, 191, 199, 204, 239, 262, 292, 296, 312, 327, 375, 376, 408, 499}, 501, -1, false},
		{[]int{47, 128, 153, 170, 305, 320, 359, 424, 458, 689, 749, 780, 830, 877, 896, 901, 989}, 31, 0, true},
	}
	for _, test := range tests {
		idx, suc := FindFirstGE(test.arr, test.target)
		if idx != test.idx || suc != test.suc {
			t.Errorf("FindLastLE(%v, %d): got(%d,%t), except(%d,%t)", test.arr, test.target, idx, suc, test.idx, test.suc)
		}
	}
}

func TestFindLastLT(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		idx    int
		suc    bool
	}{
		{[]int{63, 161, 171, 376, 472}, 240, 2, true},
		{[]int{12, 34, 118, 168, 214, 243, 314, 403, 485, 493, 493}, 214, 4, true},
		{[]int{88, 140, 144, 160, 174, 179, 242, 248, 369, 457, 467}, 113, 0, true},
		{[]int{33, 162, 192, 245, 276, 289, 430, 436}, 1306, 7, true},
		{[]int{47, 128, 153, 170, 305, 320, 359, 424, 458, 689, 749, 780, 830, 877, 896, 901, 989}, 31, -1, false},
	}
	for _, test := range tests {
		idx, suc := FindLastLE(test.arr, test.target)
		if idx != test.idx || suc != test.suc {
			t.Errorf("FindLastLE(%v, %d): got(%d,%t), except(%d,%t)", test.arr, test.target, idx, suc, test.idx, test.suc)
		}
	}
}

func TestFindFirstGT(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		idx    int
		suc    bool
	}{
		{[]int{35, 93, 163, 276, 333, 347, 415, 476}, 207, 3, true},
		{[]int{50, 66, 74, 135, 147, 158, 163, 196, 199, 223, 291, 293, 307, 349, 363, 389, 419, 444, 471, 493}, 531, -1, false},
		{[]int{39, 53, 76, 115, 155, 182, 182, 192, 196, 261, 276, 386, 410, 428, 438, 470, 479}, 192, 8, true},
		{[]int{31, 130, 135, 139, 191, 199, 204, 239, 262, 292, 296, 312, 327, 375, 376, 408, 499}, 501, -1, false},
		{[]int{47, 128, 153, 170, 305, 320, 359, 424, 458, 689, 749, 780, 830, 877, 896, 901, 989}, 31, 0, true},
	}
	for _, test := range tests {
		idx, suc := FindFirstGT(test.arr, test.target)
		if idx != test.idx || suc != test.suc {
			t.Errorf("FindLastLE(%v, %d): got(%d,%t), except(%d,%t)", test.arr, test.target, idx, suc, test.idx, test.suc)
		}
	}
}
