package leetcode

import "sort"

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(xs []int, x int) []int {
	i := sort.SearchInts(xs, x)
	if i >= len(xs) || xs[i] != x {
		return []int{-1, -1}
	}

	j := sort.SearchInts(xs, x+1)
	return []int{i, j - 1}
}
