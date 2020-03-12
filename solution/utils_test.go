package solution

import "testing"

func TestGcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{8, 6, 2},
		{5, 3, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := gcd(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
