package solution

func findPeakElement(nums []int) int {
	N := len(nums)

	for i := 0; i < N-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}

	return N - 1
}
