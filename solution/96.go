package solution

var memo = make(map[int]int)

func numTrees(n int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n <= 1 {
		return 1
	}
	count := 0
	for i := 1; i <= n; i++ {
		count += numTrees(i-1) * numTrees(n-i)
	}
	memo[n] = count
	return count
}
