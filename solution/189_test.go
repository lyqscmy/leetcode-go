package solution

import "testing"

func Test_rotate2(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rotate2(tt.nums, tt.k)
		})
	}
}
