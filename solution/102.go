package solution

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	path := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	tmp := make([]*TreeNode, 0)
	for len(q) > 0 {
		for i := 0; i < len(q); i++ {
			n := q[i]
			path = append(path, n.Val)
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		res = append(res, clone(path))
		path = path[:0]
		q, tmp = tmp, q
		tmp = tmp[:0]
	}
	return res
}
