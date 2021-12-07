package solution

import (
	"testing"
)

func Test_colorBorder(t *testing.T) {
	tests := []struct {
		grid  [][]int
		row   int
		col   int
		color int
	}{
		// {[][]int{{1, 1}, {1, 2}}, 0, 0, 3},
		{[][]int{{1, 2, 1, 2, 1, 2}, {2, 2, 2, 2, 1, 2}, {1, 2, 2, 2, 1, 2}}, 1, 3, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := colorBorder(tt.grid, tt.row, tt.col, tt.color)
			t.Error(got)
		})
	}
}
