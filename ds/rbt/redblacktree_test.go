package rbt

import (
	"math/rand"
	"testing"
	"time"
)

var (
	rng      = rand.New(rand.NewSource(time.Now().Unix()))
	tree     = New()
	N    int = 1e3
	min  int = 0
	max  int = 1e5
)

type key int

func (v key) CompareTo(i Comparator) int {
	if v < i.(key) {
		return -1
	}
	if v > i.(key) {
		return 1
	}
	return 0
}

func randKey() Comparator {
	return key(rng.Intn(max-min-1) + min + 1)
}

func TestRedBlackBST(t *testing.T) {
	m := make(map[Comparator]interface{})
	// put
	for i := 0; i < N; i++ {
		k, v := randKey(), i
		_ = tree.Put(k, v)
		m[k] = v
		verify(t)
	}
	// get
	for i := 0; i < N; i++ {
		k := randKey()
		v, _ := tree.Get(k)
		compareVal(t, v, m[k])
	}
	// delete
	for i := 0; i < N; i++ {
		k := randKey()
		_ = tree.Delete(k)
		verify(t)
	}
}

func compareVal(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Error("The value found does not match")
	}
}

func verify(t *testing.T) {
	if !tree.isSizeConsistent() {
		t.Error("Subtree counts not consistent")
	}
	// test BST
	if !tree.isBST() {
		t.Error("Not in symmetric order")
	}
	// test RBT
	if !tree.isRootBlack() {
		t.Error("Root is not black")
	}
	if !tree.is23() {
		t.Error("Not a 2-3 tree")
	}
	if !tree.isBalanced() {
		t.Error("Not balanced")
	}
}

func (t *RedBlackBST) isBST() bool {
	return t.root.isBST(key(min), key(max))
}

func (n *node) isBST(min, max Comparator) bool {
	if n == nil {
		return true
	}
	if n.key.CompareTo(min) <= 0 {
		return false
	}
	if n.key.CompareTo(max) >= 0 {
		return false
	}
	return n.left.isBST(min, n.key) && n.right.isBST(n.key, max)
}

func (t *RedBlackBST) isSizeConsistent() bool {
	return t.root.isSizeConsistent()
}

func (n *node) isSizeConsistent() bool {
	if n == nil {
		return true
	}
	if n.size != n.left.Size()+n.right.Size()+1 {
		return false
	}
	return n.left.isSizeConsistent() && n.right.isSizeConsistent()
}

func (t *RedBlackBST) isRootBlack() bool {
	return !t.root.isRed()
}

func (t *RedBlackBST) is23() bool {
	return t.root.is23()
}

func (n *node) is23() bool {
	if n == nil {
		return true
	}
	if n.right.isRed() {
		return false
	}
	if n.isRed() && n.left.isRed() {
		return false
	}
	return n.left.is23() && n.right.is23()
}

func (t *RedBlackBST) isBalanced() bool {
	black := 0
	x := t.root
	for x != nil {
		if !x.isRed() {
			black++
		}
		x = x.left
	}
	return t.root.isBalanced(black)
}

func (n *node) isBalanced(black int) bool {
	if n == nil {
		return black == 0
	}
	if !n.isRed() {
		black--
	}
	return n.left.isBalanced(black) && n.right.isBalanced(black)
}
