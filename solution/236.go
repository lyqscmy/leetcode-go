package solution

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	postorder(root, p, q, func(node *TreeNode) { res = node })
	return res
}

func postorder(root, p, q *TreeNode, f func(node *TreeNode)) bool {
	if root == nil {
		return false
	}

	left := postorder(root.Left, p, q, f)
	right := postorder(root.Right, p, q, f)
	mid := false
	if root == p || root == q {
		mid = true
	}
	if (mid && left) || (mid && right) || (left && right) {
		f(root)
	}

	return mid || left || right
}
