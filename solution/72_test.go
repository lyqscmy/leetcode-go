package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"a", "a", 0},
		{"a", "ab", 1},
		{"", "", 0},
		{"sea", "ate", 3},
		{"kitten", "sitting", 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, tt.want, minDistance(tt.word1, tt.word2), "minDistance(%v, %v)", tt.word1, tt.word2)
		})
	}
}
