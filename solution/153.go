package solution

func findMin(nums []int) int {
	pivot := findPivot(nums)
	return nums[pivot]
}
