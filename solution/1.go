package solution

import "sort"

// https://leetcode.com/problems/two-sum/
func twoSum(xs []int, target int) []int {
	res := make([]int, 0, 2)
	// x -> (index, y)
	m := make(map[int]int, len(xs))
	for i, x := range xs {
		y := target - x
		if j, ok := m[y]; ok {
			res = append(res, i, j)
		}
		m[x] = i
	}
	sort.Ints(res)
	return res
}
