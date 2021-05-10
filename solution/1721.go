package solution

func swapNodes(head *ListNode, k int) *ListNode {
	p := nth(head, k-1)
	n := length(head)
	q := nth(head, n-k)
	p.Val, q.Val = q.Val, p.Val
	return head
}
