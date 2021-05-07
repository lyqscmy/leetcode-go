package solution

func findTargetSumWays(nums []int, S int) int {
	pre := make(map[int]int)
	cur := make(map[int]int)
	pre[nums[0]] += 1
	pre[-nums[0]] += 1
	for i := 1; i < len(nums); i++ {
		for k, v := range pre {
			cur[k+nums[i]] += v
			cur[k-nums[i]] += v
		}
		pre, cur = cur, pre
		resetMap(cur)
	}

	return pre[S]
}

func resetMap(m map[int]int) {
	for k := range m {
		delete(m, k)
	}
}
