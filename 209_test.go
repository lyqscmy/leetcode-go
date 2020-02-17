package leetcode

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		s    int
		nums []int
		want int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minSubArrayLen(tt.s, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
