package leetcode

func rightSideView(root *TreeNode) []int {
	q1 := NewDeque()
	q2 := NewDeque()
	q.push(root)
	res := make([]int, 0, 4)
	// 队列中保存了一层的元素
	for !q1.empty() {
		var node *TreeNode
		for !q1.empty() {
			node = q1.pop()
			if node.Left!=nil {
				q2.push(node.Left)
			}
			if node.Right!=nil {
				q2.push(node.Right)
			}
		}
		res = append(res, node.Val)
		q1, q2 = q2, q1
	}
	return q1.items()
}
