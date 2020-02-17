package math

/*
T(n) = O(mn)
Solve with queue.
*/
func JosephusQueue(n, m int) int {
	queue := make([]int, n)
	for i := 1; i <= n; i++ {
		queue[i-1] = i
	}
	for len(queue) != 1 {
		for i := 1; i < m; i++ {
			el := queue[0]
			queue = queue[1:]
			queue = append(queue, el)
		}
		queue = queue[1:]
	}
	return queue[0]
}

/*
T(n) = O(n)
f(n,m) = 0, n=1
f(n,m) = [f(n-1,m)+m]%n, n>1
*/
func JosephusForm1(n, m int) int {
	var last int
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last + 1
}

/*
T(n) = O(logn)
current: N = n + k(m-1) + d
previous: k + N - n
K = (N - n - 1) / (m - 1)
The solution is as follows:
N = mn
while N > n:
	N = k + N - n
return N

Let D = mn + 1 - N = math.Ceil((m*D/(m-1)))
D = 1
while D <= (m-1)*n:
	D = math.Ceil((m*D/(m-1)))
return mn - D + 1
*/
func JosephusForm2(n, m int) int {
	ceil := func(x, y int) int {
		if x%y != 0 {
			return x/y + 1
		}
		return x / y
	}
	D := 1
	end := (m - 1) * n
	for D <= end {
		D = ceil(m*D, m-1)
	}
	return m*n + 1 - D
}
