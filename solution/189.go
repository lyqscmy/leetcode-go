package solution

func rotate2(nums []int, k int) {
	k = k % len(nums)
	tmp := make([]int, k)
	p := len(nums) - k
	copy(tmp, nums[p:])
	copy(nums[k:], nums[:p])
	copy(nums[:k], tmp)
}
