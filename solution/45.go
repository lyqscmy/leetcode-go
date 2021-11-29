package solution

func jump(nums []int) int {
	const intSize = 32 << (^uint(0) >> 63) // 32 or 64
	const MaxInt = 1<<(intSize-1) - 1

	N := len(nums)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = MaxInt
	}
	dp[0] = 0
	for i := 0; i < N-1; i++ {
		num := nums[i]
		cur := dp[i] + 1
		for j := 1; j <= num; j++ {
			if i+j < N && cur < dp[i+j] {
				dp[i+j] = cur
			}
		}
	}
	return dp[N-1]
}
