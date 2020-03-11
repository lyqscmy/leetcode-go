package leetcode

func sum(xs []int) int {
	acc := 0
	for _, x := range xs {
		acc += x
	}
	return acc
}
