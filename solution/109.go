package solution

func sortedListToBST(head *ListNode) *TreeNode {
	xs := make([]int, 0)
	for head != nil {
		xs = append(xs, head.Val)
		head = head.Next
	}
	return sliceToBST(xs)
}

func sliceToBST(xs []int) *TreeNode {
	if len(xs) == 0 {
		return nil
	}
	mid := len(xs) / 2
	return &TreeNode{xs[mid], sliceToBST(xs[:mid]), sliceToBST(xs[mid+1:])}
}
