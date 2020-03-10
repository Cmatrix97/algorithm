package power

import "testing"

func TestFindPowerLE(t *testing.T)  {
	var tests = []struct {
		input int
		want int
	} {
		{1, 1},
		{2, 2},
		{3, 2},
		{7, 4},
		{50, 32},
		{63, 32},
		{64, 64},
		{65, 64},
		{314159265358, 274877906944},
	}
	for _, test := range tests {
		if got := FindPowerLE(test.input); got != test.want {
			t.Errorf("findPowerLE(%d) = %d, except %d", test.input, got, test.want)
		}
	}
}