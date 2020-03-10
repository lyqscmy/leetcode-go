package leetcode

// https://leetcode.com/problems/first-missing-positive/
func firstMissingPositive(xs []int) int {
	N := len(xs)
	bm := make([]bool, len(xs)+1)
	for _, x := range xs {
		if x >= 1 && x <= N {
			bm[x-1] = true
		}
	}

	for k, v := range bm {
		if !v {
			return k + 1
		}
	}
	return 1
}
