package solution

import (
	"testing"
)

func Test_exist(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjacent(t *testing.T) {
	tests := []struct {
		name string
		a    Point
		b    Point
		want bool
	}{
		{"", Point{0, 0}, Point{0, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacent(tt.a, tt.b); got != tt.want {
				t.Errorf("adjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
