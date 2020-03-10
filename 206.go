package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	res, _ := reverseList2Impl(head)
	return res
}

func reverseList2Impl(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	h, tail := reverseList2Impl(head.Next)
	tail.Next = head
	head.Next = nil
	return h, head
}
