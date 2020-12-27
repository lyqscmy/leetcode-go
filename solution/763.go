package solution

func partitionLabels(S string) []int {
	last := make([]int, 26)
	for i := 0; i < len(S); i++ {
		last[S[i]-'a'] = i
	}

	anchor := 0
	j := 0
	res := make([]int, 0)
	for i := 0; i < len(S); i++ {
		j = max(j, last[S[i]-'a'])
		if i == j {
			res = append(res, j-anchor+1)
			anchor = i + 1
		}
	}
	return res
}
