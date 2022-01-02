package solution

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return nil
	}
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		base := i * n
		for j := 0; j < n; j++ {
			matrix[i][j] = original[base+j]
		}
	}
	return matrix
}
