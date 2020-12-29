package solution

func productExceptSelf(nums []int) []int {
	N := len(nums)
	res := make([]int, N)
	res[0] = 1
	for i := 1; i < N; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	right := 1
	for i := N - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right *= nums[i]
	}
	return res
}
