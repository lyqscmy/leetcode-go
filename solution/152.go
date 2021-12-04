package solution

func maxProduct(nums []int) int {
	type tuple struct {
		pos int
		neg int
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]tuple, len(nums)+1)
	dp[0] = tuple{0, 0}
	for i, num := range nums {
		t := tuple{num, num}
		if num >= 0 {
			t.neg *= dp[i].neg
			if dp[i].pos > 0 {
				t.pos *= dp[i].pos
			}
		} else {
			t.pos *= dp[i].neg
			if dp[i].pos > 0 {
				t.neg *= dp[i].pos
			}
		}
		dp[i+1] = t
	}

	res := 0
	for _, t := range dp {
		if t.pos > res {
			res = t.pos
		}
	}
	return res
}
