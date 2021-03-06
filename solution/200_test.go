package solution

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{
			{0x31, 0x31, 0x31, 0x31, 0x30},
			{0x31, 0x31, 0x30, 0x31, 0x30},
			{0x31, 0x31, 0x30, 0x30, 0x30},
			{0x30, 0x30, 0x30, 0x30, 0x30},
		},
			1},
		{
			[][]byte{
				{0x31, 0x31, 0x30, 0x30, 0x30},
				{0x31, 0x31, 0x30, 0x30, 0x30},
				{0x30, 0x30, 0x31, 0x30, 0x30},
				{0x30, 0x30, 0x30, 0x31, 0x31},
			}, 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
