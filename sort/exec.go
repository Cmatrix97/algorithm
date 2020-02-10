package sort

import (
	"math/rand"
)

func Sort(strategy func(a []int)) {
	rand.Seed(314159265357)
	n := 10000
	a := rand.Perm(n)
	strategy(a)
}
