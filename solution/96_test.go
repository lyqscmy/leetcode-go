package solution

import (
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
