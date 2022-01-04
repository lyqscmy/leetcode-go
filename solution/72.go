package solution

func minDistance(word1 string, word2 string) int {
	M := len(word1)
	N := len(word2)
	row1 := make([]int, M+1)
	row2 := make([]int, M+1)
	for i := 0; i <= M; i++ {
		row1[i] = i
	}
	col := make([]int, N+1)
	for i := 0; i <= N; i++ {
		col[i] = i
	}
	for i := 1; i <= N; i++ {
		row2[0] = col[i]
		for j := 1; j <= M; j++ {
			subCost := 0
			if word1[j-1] != word2[i-1] {
				subCost = 1
			}
			row2[j] = min(row1[j]+1, row2[j-1]+1, row1[j-1]+subCost)
		}
		row1, row2 = row2, row1
	}
	return row1[M]
}
