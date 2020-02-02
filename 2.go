package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l3 := dummy
	carry := 0
	for l1 !=nil || l2 !=nil {
		if l1!=nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2!=nil {
			carry += l2.Val
			l2 = l2.Next
		}

		l3.Next = &ListNode{Val:carry%10}
		l3=l3.Next
		carry /= 10
	}

	if carry != 0 {
		l3.Next = &ListNode{Val:carry}
	}

	return dummy.Next
}