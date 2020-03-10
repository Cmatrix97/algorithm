package power

// findPowerLE returns the largest power exponent less than N.
// eg. 10001001 -> 11001001 -> 11111001 -> 11111111 -> 100000000 -> 10000000
func FindPowerLE(N int) int {
	for i := 0; i <= 5; i++ {
		N |= N >> (1 << i)
	}
	return (N + 1) >> 1
}
