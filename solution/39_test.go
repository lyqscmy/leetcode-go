package solution

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{7}, 7, [][]int{{7}}},
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := combinationSum(tt.candidates, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
