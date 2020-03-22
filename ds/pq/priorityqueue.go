package pq

import "errors"

// priorityQueue is a heap implements by slice.
// None of method are thread-safe.
type priorityQueue struct {
	arr []int // implements heap by array
	n   int   // number of priority queue members
	// thread-safe
	// mu sync.Mutex
}

// ErrIndexOutOfBounds is returned when an out-of-bounds index is accessed.
var ErrIndexOutOfBounds = errors.New("The index accessed is out of bounds.")

// New creates and returns values of the type priorityQueue.
func New() *priorityQueue {
	pq := new(priorityQueue)
	pq.arr = make([]int, 1)
	return pq
}

// Empty returns whether pq is empty.
func (pq priorityQueue) Empty() bool {
	return pq.n == 0
}

// Size returns number of pq members.
func (pq priorityQueue) Size() int {
	return pq.n
}

// Push adds element v to the pq.
func (pq *priorityQueue) Push(v int) {
	pq.arr = append(pq.arr, v)
	pq.n++
	pq.swim(pq.n)
}

// Pop returns and removes the root element of pq.
func (pq *priorityQueue) Pop() (int, error) {
	if pq.n == 0 {
		return 0, ErrIndexOutOfBounds
	}
	root := pq.arr[1]
	pq.exch(1, pq.n)
	pq.arr = pq.arr[:pq.n]
	pq.n--
	pq.sink(1)
	return root, nil
}

// Top retures the root element of pq.
func (pq priorityQueue) Top() (int, error) {
	if pq.n == 0 {
		return 0, ErrIndexOutOfBounds
	}
	return pq.arr[1], nil
}

// TopK returns the first k elements.
func (pq priorityQueue) TopK(k int) ([]int, error) {
	if k <= 0 || k > pq.n {
		return nil, ErrIndexOutOfBounds
	}
	return pq.arr[1 : 1+k], nil
}

func (pq *priorityQueue) comp(i, j int) bool {
	// maxPQ
	return pq.arr[i] < pq.arr[j]
	// minPQ
	// return pq.arr[i] > pq.arr[j]
}

func (pq *priorityQueue) exch(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq *priorityQueue) swim(k int) {
	for k > 1 && pq.comp(k>>1, k) {
		pq.exch(k>>1, k)
		k >>= 1
	}
}

func (pq *priorityQueue) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.comp(j, j+1) {
			j++
		}
		if !pq.comp(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}
