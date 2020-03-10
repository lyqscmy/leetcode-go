package leetcode

import (
	"testing"
)

func TestRverseList2(t *testing.T) {
	head := &ListNode{5, nil}
	head = &ListNode{4, head}
	head = &ListNode{3, head}
	head = &ListNode{2, head}
	head = &ListNode{1, head}

	want := &ListNode{1, nil}
	want = &ListNode{2, want}
	want = &ListNode{3, want}
	want = &ListNode{4, want}
	want = &ListNode{5, want}
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{head, want},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseList2(tt.head); !got.equal(tt.want) {
				t.Errorf("reverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
