package solution

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	res := make([]int, 0)
	nextLevel := make([]*TreeNode, 0)
	for len(curLevel) > 0 {
		var node *TreeNode
		for len(curLevel) > 0 {
			node = curLevel[0]
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
			curLevel = curLevel[1:]
		}
		res = append(res, node.Val)
		curLevel, nextLevel = nextLevel, curLevel
	}
	return res
}
