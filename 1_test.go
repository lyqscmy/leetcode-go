package leetcode

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		xs   []int
		x    int
		want []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoSum(tt.xs, tt.x); !Compare(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
