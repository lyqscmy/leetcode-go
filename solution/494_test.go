package solution

import (
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		nums []int
		S    int
		want int
	}{
		//{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1, 256},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findTargetSumWays(tt.nums, tt.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
