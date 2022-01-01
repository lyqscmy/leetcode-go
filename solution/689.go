package solution

import "sort"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	dp := make([]Tuple, 0) // tuple.a sum, tuple.b index
	s := 0
	for i := 0; i < k; i++ {
		s += nums[i]
	}
	first := nums[0]
	dp = append(dp, Tuple{s, 0})
	for i := 1; i < len(nums)-k+1; i++ {
		s = s - first + nums[i+k-1]
		first = nums[i]
		dp = append(dp, Tuple{s, i})
	}
	sort.Slice(dp, func(i, j int) bool {
		a, b := dp[i], dp[j]
		if a.x > b.x {
			return true
		}
		if a.x < b.x {
			return false
		}
		return a.y < b.y
	})

	max := MinInt
	var A, B, C int
	for i := 0; i < len(dp)-2; i++ {
		var a, b, c Tuple
		a = dp[i]
		for j := i + 1; j < len(dp)-1; j++ {
			b = dp[j]
			if abs(b.y-a.y) < k {
				continue
			}
			for r := j + 1; r < len(dp); r++ {
				c = dp[r]
				if abs(c.y-b.y) < k || abs(c.y-a.y) < k {
					continue
				}
				s := a.x + b.x + c.x
				if s > max {
					max = s
					A, B, C = a.y, b.y, c.y
				}
			}
		}
	}
	res := []int{A, B, C}
	sort.Ints(res)
	return res
}
