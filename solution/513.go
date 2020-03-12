package solution

// res 记录从上到下每一层最左的值
// level 表示当前的层数
// 层序遍历树，从上到下从左到右，同一层的节点，只有最左的节点能加入res
func findLeftMostBottomNode(root *TreeNode) int {
	res := []int{-1}
	res = root.dfs(1, res)
	return res[len(res)-1]
}

func (root *TreeNode) dfs(index int, res []int) []int {
	if root == nil {
		return res
	}

	if index > len(res) {
		res = append(res, root.Val)
	}

	res = root.Left.dfs(index+1, res)
	res = root.Right.dfs(index+1, res)
	return res
}
