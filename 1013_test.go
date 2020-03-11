package leetcode

import "testing"

func TestCanThreePartsEqualSum(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}, true},
		{[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}, false},
		{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}, true},
		{[]int{18, 12, -18, 18, -19, -1, 10, 10}, true},
		{[]int{1, -1, 1, -1}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
