package solution

// len(nums) > 0
func rob(nums []int) int {
	N := len(nums)
	dp := make([]int, N+2)
	max := 0
	for i := 0; i < N; i++ {
		j := i + 2
		dp[j] = maxInt(dp[j-2]+nums[i], dp[j-1])
		if dp[j] > max {
			max = dp[j]
		}
		// fmt.Println(dp)
	}

	return max
}
