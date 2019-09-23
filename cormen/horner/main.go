package main

//index 0 is coeffecient for a0, index 1 is coeefficent for a1.. etc
//given a polynomial in the form ... x^3*a3 + x^2*a2 + x*a1 + a0
//for example 2x^3 - 6x^2 + 2x - 1
//example := poly([]int{-1, 2, -6, 2})
type poly []int

func main() {

}

func simple(x int, p poly) int {
	y := 0
	for i := len(p) - 1; i > -1; i-- {
		if i == 0 { // x^0 is the same a 1 (ie 1*p[i])
			y = y + p[i]
			continue
		}
		z := x
		for j := 1; j < i; j++ {
			z = z * x
		}
		y = y + z*p[i]
	}
	return y
}

func horner(x int, p poly) int {
	y := 0
	for i := len(p) - 1; i > -1; i-- {
		y = p[i] + x*y
	}
	return y
}
