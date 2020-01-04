package leetcode

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		xs   []int
		x    int
		want []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := searchRange(tt.xs, tt.x); !Compare(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
