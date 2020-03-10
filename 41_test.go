package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{}, 1},
		{[]int{1}, 2},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := firstMissingPositive(tt.xs); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
