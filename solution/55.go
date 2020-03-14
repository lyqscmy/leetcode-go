package solution

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	N := nums[0]
	if N == 0 {
		return false
	}
	for i := N; i > 0; i++ {
		if canJump(nums[i:]) {
			return true
		}
	}
	return false
}
