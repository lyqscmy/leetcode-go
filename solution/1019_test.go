package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	cases := []struct {
		xs   []int
		want []int
	}{
		{[]int{2, 1, 5}, []int{5, 5, 0}},
		{[]int{2, 7, 4, 3, 5}, []int{7, 0, 5, 5, 0}},
		{[]int{1, 7, 5, 1, 9, 2, 5, 1}, []int{7, 9, 9, 9, 0, 5, 0, 0}},
	}
	for _, c := range cases {
		head := NewListNode(c.xs)
		got := nextLargerNodes(head)
		assert.Equal(t, got, c.want)
	}
}
