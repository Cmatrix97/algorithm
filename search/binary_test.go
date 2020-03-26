package search

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	maxLen    = 100   // length of test array
	maxVal    = 200   // max value in array
	testTimes = 10000 // test times
)

var rng *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	rng = rand.New(rand.NewSource(seed))
}

func randomArray(rng *rand.Rand) (arr []int, target int) {
	n := rng.Intn(maxLen)
	arr = make([]int, n)
	for i := range arr {
		arr[i] = rng.Intn(maxVal)
	}
	sort.Ints(arr)
	target = rand.Intn(maxLen)
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

func linearLowerBound(arr []int, target int) int {
	for i := range arr {
		if arr[i] >= target {
			return i
		}
	}
	return len(arr)
}

func linearUpperBound(arr []int, target int) int {
	for i := range arr {
		if arr[i] > target {
			return i
		}
	}
	return len(arr)
}

func TestBinarySearch(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdxs, wantSuc := linearSearch(arr, target)
		idx, suc := BinarySearch(arr, 0, len(arr)-1, target)
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

func TestLowerBound(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearLowerBound(arr, target)
		idx := LowerBound(arr, 0, len(arr), target)
		if idx != wantIdx {
			t.Errorf("LowerBound(%v, %d): got %d, want %d", arr, target, idx, wantIdx)
		}
	}
}

func TestUpperBound(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearUpperBound(arr, target)
		idx := UpperBound(arr, 0, len(arr), target)
		if idx != wantIdx {
			t.Errorf("UpperBound(%v, %d): got %d, want %d", arr, target, idx, wantIdx)
		}
	}
}

func TestLastLE(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearUpperBound(arr, target) - 1
		idx, suc := LastLE(arr, target)
		if suc != (wantIdx != -1) {
			t.Errorf("LastLE(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
		if suc && idx != wantIdx {
			t.Errorf("LastLE(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
	}
}

func TestFirstGE(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearLowerBound(arr, target)
		idx, suc := FirstGE(arr, target)
		if suc != (wantIdx != len(arr)) {
			t.Errorf("FirstGE(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
		if suc && idx != wantIdx {
			t.Errorf("FirstGE(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
	}
}

func TestLastLT(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearLowerBound(arr, target) - 1
		idx, suc := LastLT(arr, target)
		if suc != (wantIdx != -1) {
			t.Errorf("LastLT(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
		if suc && idx != wantIdx {
			t.Errorf("LastLT(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
	}
}

func TestFirstGT(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		arr, target := randomArray(rng)
		wantIdx := linearUpperBound(arr, target)
		idx, suc := FirstGT(arr, target)
		if suc != (wantIdx != len(arr)) {
			t.Errorf("FirstGT(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
		if suc && idx != wantIdx {
			t.Errorf("FirstGT(%v, %d): got(%d,%t), want(%d)", arr, target, idx, suc, wantIdx)
		}
	}
}
