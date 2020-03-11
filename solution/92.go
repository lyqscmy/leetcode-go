package solution

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := m; i < n; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}
