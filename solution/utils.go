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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxOfInts(xs []int) int {
	max := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max
}
