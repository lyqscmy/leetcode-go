package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) equal(other *TreeNode) bool {
	return false
}

func (root *TreeNode) findLeftMostBottomNode() int {
	res := []int{-1}
	res = root.findLeftMostBottomNodeImpl(1, res)
	return res[len(res)-1]
}

func (root *TreeNode) findLeftMostBottomNodeImpl(index int, res []int) []int {
	if root == nil {
		return res
	}

	if index > len(res) {
		res = append(res, root.Val)
	}

	res = root.Left.findLeftMostBottomNodeImpl(index+1, res)
	res = root.Right.findLeftMostBottomNodeImpl(index+1, res)
	return res
}
