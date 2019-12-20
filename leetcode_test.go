package leetcode

import (
	"testing"
)

func TestFindLeftMostBottomNode(t *testing.T) {
	root := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	got := root.findLeftMostBottomNode()
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
