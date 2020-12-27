package solution

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := permute(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
