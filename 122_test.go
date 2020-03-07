package leetcode

import "testing"

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxProfit2(tt.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
