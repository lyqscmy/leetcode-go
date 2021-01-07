package solution

import (
	"testing"
)

func Test_maxProfit3(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxProfit3(tt.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
		})
	}
}
