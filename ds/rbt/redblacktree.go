package rbt

import "errors"

type color bool

const (
	Red   = true
	Black = false
)

// Comparator ...
type Comparator interface {
	CompareTo(Comparator) int
}

type node struct {
	key   Comparator  // key
	val   interface{} // associated data
	left  *node       // links to left asubtrees
	right *node       // links to right subtrees
	color color       // color of parent link
	size  int         // subtree count
}

// NewNode ...
func NewNode(key Comparator, val interface{}, color color, size int) *node {
	return &node{
		key:   key,
		val:   val,
		color: color,
		size:  size,
	}
}

// RedBlackBST ...
type RedBlackBST struct {
	root *node // root of the RBT
}

// New creates and returns a red-black tree.
func New() *RedBlackBST {
	return new(RedBlackBST)
}

func (t *RedBlackBST) isEmpty() bool {
	return t.root == nil
}

// Size returns the number of key-value pairs in this red-black tree.
func (t *RedBlackBST) Size() int {
	return t.root.Size()
}

// Size returns the number of key-value pairs in a subtree rooted at n.
func (n *node) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color == Red
}

var ErrUnderflow = errors.New("BST underflow")
var ErrArgumentNil = errors.New("argument to delete is nil")

// Get returns the value associated with the given key.
func (t *RedBlackBST) Get(key Comparator) (interface{}, error) {
	if key == nil {
		return nil, ErrArgumentNil
	}
	return t.root.get(key), nil
}

func (n *node) get(key Comparator) interface{} {
	for n != nil {
		if cmp := key.CompareTo(n.key); cmp < 0 {
			n = n.left
		} else if cmp > 0 {
			n = n.right
		} else {
			return n.val
		}
	}
	return nil
}

// Put Inserts the specified key-value pair into the red-black tree, overwriting the old
// value with the new value if the red-black tree already contains the specified key.
// Deletes the specified key (and its associated value) from this red-black tree
// if the specified value is nil.
func (t *RedBlackBST) Put(key Comparator, val interface{}) error {
	if key == nil {
		return ErrArgumentNil
	}
	if val == nil {
		return t.Delete(key)
	}
	t.root = t.root.put(key, val)
	t.root.color = Black
	return nil
}

// insert the key-value pair in the subtree rooted at n
func (n *node) put(key Comparator, val interface{}) *node {
	if n == nil {
		return NewNode(key, val, Red, 1)
	}

	if cmp := key.CompareTo(n.key); cmp < 0 {
		n.left = n.left.put(key, val)
	} else if cmp > 0 {
		n.right = n.right.put(key, val)
	} else {
		n.val = val
	}

	if !n.left.isRed() && n.right.isRed() {
		n = n.rotateLeft()
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = n.rotateRight()
	}
	if n.left.isRed() && n.right.isRed() {
		n.flipColors()
	}
	n.size = n.left.Size() + n.right.Size() + 1
	return n
}

// make a left-leaning link lean to the right
func (n *node) rotateRight() *node {
	x := n.left
	x.right, n.left = n, x.right
	x.color, n.color = n.color, x.color
	x.size, n.size = n.Size(), n.left.Size()+n.right.Size()+1
	return x
}

// make a right-leaning link lean to the left
func (n *node) rotateLeft() *node {
	x := n.right
	x.left, n.right = n, x.left
	x.color, n.color = n.color, x.color
	x.size, n.size = n.Size(), n.left.Size()+n.right.Size()+1
	return x
}

// flip the colors of a node and its two children
func (n *node) flipColors() {
	n.color, n.left.color, n.right.color = !n.color, !n.left.color, !n.right.color
}

// DeleteMin removes the smallest key and associated value from the red-black tree.
func (t *RedBlackBST) DeleteMin() error {
	if t.isEmpty() {
		return ErrUnderflow
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = Red
	}
	t.root = t.root.deleteMin()
	if !t.isEmpty() {
		t.root.color = Black
	}
	return nil
}

// delete the key-value pair with the minimum key rooted at n
func (n *node) deleteMin() *node {
	if n.left == nil {
		return nil
	}
	if !n.left.isRed() && !n.left.left.isRed() {
		n = n.moveRedLeft()
	}
	n.left = n.left.deleteMin()
	return n.balance()
}

// DeleteMax removes the largest key and associated value from the red-black tree.
func (t *RedBlackBST) DeleteMax() error {
	if t.isEmpty() {
		return ErrUnderflow
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = Red
	}
	t.root = t.root.deleteMax()
	if !t.isEmpty() {
		t.root.color = Black
	}
	return nil
}

// delete the key-value pair with the maximum key rooted at n
func (n *node) deleteMax() *node {
	if n.left.isRed() {
		n = n.rotateRight()
	}
	if n.right == nil {
		return nil
	}
	if !n.right.isRed() && !n.right.left.isRed() {
		n = n.moveRedRight()
	}
	n.right = n.right.deleteMax()
	return n.balance()
}

// Delete removes the specified key and its associated value from the red-black tree.
func (t *RedBlackBST) Delete(key Comparator) error {
	if key == nil {
		return ErrArgumentNil
	}
	if t.root.get(key) == nil {
		return nil
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = Red
	}
	t.root = t.root.delete(key)
	if !t.isEmpty() {
		t.root.color = Black
	}
	return nil
}

// delete the key-value pair with the given key rooted at n
func (n *node) delete(key Comparator) *node {
	if key.CompareTo(n.key) < 0 {
		if !n.left.isRed() && !n.left.left.isRed() {
			n = n.moveRedLeft()
		}
		n.left = n.left.delete(key)
	} else {
		if n.left.isRed() {
			n = n.rotateRight()
		}
		if key.CompareTo(n.key) == 0 && n.right == nil {
			return nil
		}
		if !n.right.isRed() && !n.right.left.isRed() {
			n = n.moveRedRight()
		}
		if key.CompareTo(n.key) == 0 {
			x := n.right.min()
			n.key, n.val = x.key, x.val
			n.right = n.right.deleteMin()
		} else {
			n.right = n.right.delete(key)
		}
	}
	return n.balance()
}

// Assuming that n is red and both n.left and n.left.left
// are black, make n.left or one of its children red.
func (n *node) moveRedLeft() *node {
	n.flipColors()
	if n.right.left.isRed() {
		n.right = n.right.rotateRight()
		n = n.rotateLeft()
		n.flipColors()
	}
	return n
}

// Assuming that n is red and both n.right and n.right.left
// are black, make n.right or one of its children red.
func (n *node) moveRedRight() *node {
	n.flipColors()
	if n.left.left.isRed() {
		n = n.rotateRight()
		n.flipColors()
	}
	return n
}

// restore red-black tree invariant
func (n *node) balance() *node {
	if n.right.isRed() {
		n = n.rotateLeft()
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = n.rotateRight()
	}
	if n.left.isRed() && n.right.isRed() {
		n.flipColors()
	}
	n.size = n.left.Size() + n.right.Size() + 1
	return n
}

// the smallest key in subtree rooted at n, nil if no such key
func (n *node) min() *node {
	if n.left == nil {
		return n
	}
	return n.left.min()
}
