package leetcode

// 累加所有上坡的线段
func maxProfit2(prices []int) int {
	N := len(prices)
	if N < 2 {
		return 0
	}
	res := 0
	pre := prices[0]
	for i := 0; i < N; i++ {
		cur := prices[i]
		if cur < pre {
			pre = cur
			continue
		}

		if cur > pre {
			res += cur - pre
			pre = cur
		}
	}
	return res
}
