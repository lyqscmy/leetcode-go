package solution

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	N := len(nums)
	if N < 2 {
		return N
	}
	sort.Ints(nums)
	max := 1
	count := 1
	pre := nums[0]
	for i := 1; i < N; i++ {
		cur := nums[i]
		if cur == pre || cur == pre+1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
		pre = cur
	}
	return max
}
