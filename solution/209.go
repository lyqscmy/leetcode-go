package solution

func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	left := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			res = min(res, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if res > len(nums) {
		res = 0
	}
	return res
}

func min(xs ...int) int {
	res := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < res {
			res = xs[i]
		}
	}
	return res
}
