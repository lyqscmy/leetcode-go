package solution

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	inorder(root, func(node *TreeNode) {
		if k > 0 {
			res = node.Val
			k--
		}
	})
	return res
}

func inorder(root *TreeNode, f func(node *TreeNode)) {
	if root == nil {
		return
	}

	inorder(root.Left, f)
	f(root)
	inorder(root.Right, f)
}
