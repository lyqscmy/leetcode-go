package solution

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		// {[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
