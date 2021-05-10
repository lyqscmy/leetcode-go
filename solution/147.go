package solution

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	for head != nil {
		pre := dummy
		cur := pre.Next
		for cur != nil && head.Val > cur.Val {
			pre = pre.Next
			cur = cur.Next
		}
		pre.Next = head
		head = head.Next
		pre.Next.Next = cur
	}
	return dummy.Next
}
