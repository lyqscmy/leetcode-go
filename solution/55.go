package solution

func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		far := nums[i] + i
		if far > max {
			max = far
		}
	}
	return true
}
