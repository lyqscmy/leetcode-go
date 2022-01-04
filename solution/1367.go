package solution

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return dfs1367(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs1367(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}
	return dfs1367(head.Next, root.Left) || dfs1367(head.Next, root.Right)
}
