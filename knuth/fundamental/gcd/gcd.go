package gcd

//E is euclids algorithm for gcd as on knuth, vol1, 3rd ed, p2
//Algorithm 1.1E. Given two positive integers m and n, find their
//greatest common divisor, that is, the largest postive integer that,
//evenly divides both m and n.
func E(m int, n int) int {
	r := m % n  // E1 divide m by n and let r be the remainder
	if r == 0 { //E2 if r == 0 n is the answer
		return n
	}
	return E(n, r) //E3 set m <- n, n <- r and go back to E1
}
