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

func TestHtod(t *testing.T) {
	h := []int{1, 2, 3, 4}
	want := []int{4, 2, 3, 1}
	got := htod(h)
	if Compare(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDtoh(t *testing.T) {
	want := []int{1, 2, 3, 4}
	d := []int{4, 2, 3, 1}
	got := dtoh(d)
	if Compare(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
