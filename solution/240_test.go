package solution

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := make([][]int, 0)
	row := []int{1, 4, 7, 11, 15}
	matrix = append(matrix, row)

	row = []int{2, 5, 8, 12, 19}
	matrix = append(matrix, row)

	row = []int{3, 6, 9, 16, 22}
	matrix = append(matrix, row)

	row = []int{10, 13, 14, 17, 24}
	matrix = append(matrix, row)

	row = []int{18, 21, 23, 26, 30}
	matrix = append(matrix, row)

	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{matrix, 5, true},
		{matrix, 20, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
