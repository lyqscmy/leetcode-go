package solution

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	left := &TreeNode{Val: 4}
	right := &TreeNode{Val: 5}
	left = &TreeNode{Val: 2, Left: left, Right: right}
	right = &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: left, Right: right}
	tests := []struct {
		root *TreeNode
		want int
	}{
		{root, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
