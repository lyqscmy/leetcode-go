package solution

import (
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	tests := []struct {
		original []int
		m        int
		n        int
	}{
		{[]int{1, 2, 3, 4}, 2, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			construct2DArray(tt.original, tt.m, tt.n)
		})
	}
}
