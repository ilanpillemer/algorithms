package gcd

import "testing"

func TestE(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{119, 544, 17},
		{544, 119, 17},
	}

	for _, test := range tests {
		if got := E(test.m, test.n); got != test.want {
			t.Errorf("GCD of %d, %d want %d got %d", test.m, test.m, test.want, got)
		}
	}
}

// m <- 119, n <-544
// m = cn + r
// 119 = c544 + 119 . c == 0
// m <- 544, n <- 119
// 544 = 4.119 + 68 . c == 4
// 544 - 4.119 = 68.. so any number that can divide left can divide the right
// ie any number that can divide both 119 and 68 without remainder can divide 544
// so working out a solution to gcd of 119 and 68 is the same answer to 119 and 544
// as 544 is the fraction 4 68/119
