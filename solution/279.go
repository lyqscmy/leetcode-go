package solution

var dp = make(map[int]int)

func numSquares(n int) int {
	if v, ok := dp[n]; ok {
		return v
	}
	min := n
	for i := 2; n >= i*i; i++ {
		x := numSquares(n - i*i)
		if x+1 < min {
			min = x + 1
		}
	}
	dp[n] = min
	return min
}
