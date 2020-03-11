package solution

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	L := invertTree(root.Left)
	R := invertTree(root.Right)
	root.Left, root.Right = R, L
	return root
}
