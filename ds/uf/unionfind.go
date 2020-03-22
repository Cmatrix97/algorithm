package uf

// unionFind is a data structure to solve dynamic connectivity problem efficiently.
// It implements by weighted quick-union with path compression.
type unionFind struct {
	parent []int // parent link (site indexed)
	size   []int // size of component for roots (site indexed)
	count  int   // number of components
}

// New creates and returns values of the type unionFind.
func New(n int) *unionFind {
	uf := new(unionFind)
	uf.count = n
	uf.parent = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Count returns the number of components.
func (uf *unionFind) Count() int {
	return uf.count
}

// Connected returns true if p and q are in the same component.
func (uf *unionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Union adds connection between p and q.
func (uf *unionFind) Union(p, q int) {
	pRoot, qRoot := uf.Find(p), uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.size[pRoot] < uf.size[qRoot] {
		uf.parent[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.size[pRoot] += uf.size[qRoot]
	}
	uf.count--
}

// Find returns the component identifier for p.
func (uf *unionFind) Find(p int) int {
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}
