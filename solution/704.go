package solution

func bsearch(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		// i<=h<j
		h := int(uint(i+j) >> 1)
		if nums[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	if i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}
