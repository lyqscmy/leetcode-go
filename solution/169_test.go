package solution

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
