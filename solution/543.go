package solution

// 路径经过的节点数
var _ans int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_ans = 0
	depth(root)
	return _ans - 1
}

// 以root为根的深度
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	L := depth(root.Left)
	R := depth(root.Right)
	_ans = max(_ans, L+R+1)
	return max(L, R) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
