package leetcode

// 寻找最小值和最大值
func maxProfit(prices []int) int {
	N := len(prices)
	if N < 2 {
		return 0
	}
	res := 0
	valley := prices[0]
	for i := 1; i < N; i++ {
		cur := prices[i]
		if cur <= valley {
			valley = cur
			continue
		}
		if profit := cur - valley; profit > res {
			res = profit
		}
	}
	return res
}
