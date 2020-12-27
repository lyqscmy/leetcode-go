package solution

func permute(nums []int) [][]int {
	return permuteImpl(len(nums), nums)
}

func permuteImpl(N int, nums []int) [][]int {
	if N == 1 {
		b := make([]int, len(nums))
		copy(b, nums)
		return [][]int{b}
	}
	res := permuteImpl(N-1, nums)
	for i := 0; i < N-1; i++ {
		if N%2 == 0 {
			nums[i], nums[N-1] = nums[N-1], nums[i]
		} else {
			nums[0], nums[N-1] = nums[N-1], nums[0]
		}
		res = append(res, permuteImpl(N-1, nums)...)
	}
	return res
}
