package leetcode

import (
	"testing"
)

func TestFindPivot(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4},
		{[]int{1}, 0},
		{[]int{1, 3}, 0},
		{[]int{3, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findPivot(tt.nums); got != tt.want {
				t.Errorf("findPivot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := rsearch(tt.nums, tt.target); got != tt.want {
				t.Errorf("rsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
