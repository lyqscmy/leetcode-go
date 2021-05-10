package solution

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	m := make(map[int]*ListNode)
	m[0] = dummy
	cur := head
	prefix := 0
	for cur != nil {
		prefix += cur.Val
		m[prefix] = cur
		cur = cur.Next
	}

	cur = dummy
	prefix = 0
	for cur != nil {
		prefix += cur.Val
		cur.Next = m[prefix].Next
		cur = cur.Next
	}
	return dummy.Next
}
