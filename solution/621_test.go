package solution

import (
	"testing"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks []byte
		n     int
		want  int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := leastInterval(tt.tasks, tt.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
