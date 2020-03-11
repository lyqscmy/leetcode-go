package solution

import (
	"strconv"
	"testing"
)

func contains(xs []int, x int) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func TestParseInt(t *testing.T) {
	var tests = []struct {
		in  string
		out int32
	}{
		{"0", 0},
		{"1", 1},
		{strconv.Itoa(1<<32 - 1), -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			act := ParseInt(tt.in)
			if act != tt.out {
				t.Errorf("got %v, want %v", act, tt.out)
			}
		})
	}
}
