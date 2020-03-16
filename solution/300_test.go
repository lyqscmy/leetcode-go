package solution

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
