package solution

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		S    string
		want []int
	}{
		{S: "ababcbacadefegdehijhklij", want: []int{9, 7, 8}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := partitionLabels(tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
