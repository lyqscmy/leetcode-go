package solution

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for i, j := m-1, 0; i >= 0 && j < n; {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
		} else {
			i--
		}
	}
	return false
}
