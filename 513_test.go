package leetcode

import (
	"testing"
)

func TestFindLeftMostBottomNode(t *testing.T) {
	t1 := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}

	left := &TreeNode{Val: 7}
	left = &TreeNode{Val: 5, Left: left}
	right := &TreeNode{Val: 6}
	right = &TreeNode{Val: 3, Left: left, Right: right}

	tests := []struct {
		root *TreeNode
		want int
	}{
		{t1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findLeftMostBottomNode(tt.root); got != tt.want {
				t.Errorf("findLeftMostBottomNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
