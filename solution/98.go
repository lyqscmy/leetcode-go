package solution

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	pre := MinInt
	f := func(val int) bool {
		if val <= pre {
			return false
		}
		pre = val
		return true
	}
	return preOrder(root, f)
}

func preOrder(root *TreeNode, f func(int) bool) bool {
	if root == nil {
		return true
	}

	return preOrder(root.Left, f) && f(root.Val) && preOrder(root.Right, f)
}
