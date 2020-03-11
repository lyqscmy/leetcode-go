package solution

import (
	"sort"
)

// nums no duplicate
func rsearch(nums []int, target int) int {
	pivot := findPivot(nums)
	res := sort.SearchInts(nums[:pivot], target)
	if res < len(nums) && nums[res] == target {
		return res
	}

	res = pivot + sort.SearchInts(nums[pivot:], target)
	if res < len(nums) && nums[res] == target {
		return res
	}
	return -1
}

// find index
func findPivot(nums []int) int {
	i, j := 0, len(nums)
	if nums[i] <= nums[j-1] {
		return 0
	}
	// nums[i] > nums[j]
	for i < j {
		// i <= h < j
		h := int(uint(i+j) >> 1)
		if nums[h-1] > nums[h] {
			return h
		}
		if nums[h] > nums[h+1] {
			return h + 1
		}
		if nums[i] < nums[h] {
			i = h + 1
		} else {
			j = h
		}
	}
	return 0
}
