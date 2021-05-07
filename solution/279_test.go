package solution

import (
	"testing"
)

func Test_numSquares(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{12, 3},
		{13, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
