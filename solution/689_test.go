package solution

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2, []int{0, 2, 4}},
		{[]int{17, 9, 3, 2, 7, 10, 20, 1, 13, 4, 5, 16, 4, 1, 17, 6, 4, 19, 8, 3}, 4, []int{3, 8, 14}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxSumOfThreeSubarrays(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
