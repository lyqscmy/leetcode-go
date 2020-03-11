package solution

func maxSubArray(nums []int) int {
	acc := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		acc += nums[i]
		if acc > max {
			max = acc
		}
		if acc < 0 {
			acc = 0
		}
	}

	if max == 0 && len(nums) != 0 {
		max = nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
	}
	return max
}
