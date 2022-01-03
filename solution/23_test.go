package solution

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
	}{
		{[]*ListNode{NewListNode([]int{1, 4, 5}), NewListNode([]int{1, 3, 4}), NewListNode([]int{2, 6})}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			mergeKLists(tt.lists)
		})
	}
}
