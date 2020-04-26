package bitvector

import (
	"bytes"
	"fmt"
)

// uintSize is the effective size of uint in bits, 32 or 64.
const uintSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Len returns the number of elements.
func (s *IntSet) Len() int {
	var count int
	for _, Word := range s.words {
		if Word == 0 {
			continue
		}
		Word &= Word - 1
		count++
	}
	return count
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	Word, bit := x/uintSize, uint(x%uintSize)
	return Word < len(s.words) && s.words[Word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	Word, bit := x/uintSize, uint(x%uintSize)
	for Word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[Word] |= 1 << bit
}

// AddAll adds a set of the non-negative values to the set.
func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

// Remove removes the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	Word, bit := x/uintSize, uint(x%uintSize)
	if Word >= len(s.words) {
		return
	}
	s.words[Word] &^= 1 << bit
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	cwords := make([]uint, len(s.words))
	copy(cwords, s.words)
	return &IntSet{cwords}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	sLen, tLen := len(s.words), len(t.words)
	if sLen > tLen {
		s.words = s.words[:tLen]
	}
	for i := range s.words {
		s.words[i] &= t.words[i]
	}
}

// DifferenceWith retain element in s which appears in s but not in t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		s.words[i] &^= tword
	}
}

// SymmetricDifference retain element in s which appears in s but not in t or appears in t but not in s.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

// Elems returns all elements in the set.
func (s *IntSet) Elems() []int {
	var elems []int
	for i, Word := range s.words {
		if Word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if Word&(1<<uint(j)) != 0 {
				elems = append(elems, uintSize*i+j)
			}
		}
	}
	return elems
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, Word := range s.words {
		if Word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if Word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
