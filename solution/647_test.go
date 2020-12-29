package solution

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countSubstrings(tt.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
