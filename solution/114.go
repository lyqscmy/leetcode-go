package solution

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	flatten(root.Left)
	node := root.Left
	for node.Right != nil {
		node = node.Right
	}
	node.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return
}
