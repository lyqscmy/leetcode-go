package leetcode

import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := &ListNode{1, nil}
	head = &ListNode{2, head}
	head = &ListNode{3, head}
	head = &ListNode{4, head}
	head = &ListNode{5, head}

	want := &ListNode{5, nil}
	want = &ListNode{2, want}
	want = &ListNode{3, want}
	want = &ListNode{4, want}
	want = &ListNode{1, want}

	tests := []struct {
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{head, 2, 4, want},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseBetween(tt.head, tt.m, tt.n); got.equal(tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
