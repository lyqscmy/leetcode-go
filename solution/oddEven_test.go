package solution

import (
	"reflect"
	"testing"
)

func Test_oddEven(t *testing.T) {
	tests := []struct {
		xs   []int
		want []int
	}{
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 1, 4, 5, 2}, []int{1, 3, 5, 2, 4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			oddEven(tt.xs)
			if !reflect.DeepEqual(tt.xs, tt.want) {
				t.Errorf("got %v, want %v", tt.xs, tt.want)
			}
		})
	}
}
