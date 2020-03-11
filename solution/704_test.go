package solution

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		// {[]int{}, 1, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		// {[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := bsearch(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
