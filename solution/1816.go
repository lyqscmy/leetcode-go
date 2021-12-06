package solution

func truncateSentence(s string, k int) string {
	if k <= 0 {
		return ""
	}
	space := make([]int, 0)
	N := len(s)
	for i := 0; i < N; i++ {
		if s[i] == ' ' {
			space = append(space, i)
		}
	}
	if k > len(space) {
		return s
	}
	return s[:space[k-1]]
}
