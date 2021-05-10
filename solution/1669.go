package solution

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p := list1
	for i := 0; i < a-1; i++ {
		p = p.Next
	}
	q := p
	for i := 0; i < b-a+2; i++ {
		q = q.Next
	}

	p.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = q
	return list1
}
