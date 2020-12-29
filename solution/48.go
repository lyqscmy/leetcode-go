package solution

func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < N; i++ {
		reverse(matrix[i])
	}
}

func reverse(xs []int) {
	N := len(xs)
	mid := N / 2
	for i := 0; i < mid; i++ {
		xs[i], xs[N-i-1] = xs[N-i-1], xs[i]
	}
}
