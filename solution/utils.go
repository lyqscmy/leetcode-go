package solution

func sum(xs []int) int {
	acc := 0
	for _, x := range xs {
		acc += x
	}
	return acc
}

// a > b
func gcd(a, b int) int {
	c := a % b
	if c == 0 {
		return b
	}
	return gcd(b, c)
}
