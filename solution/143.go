package solution

func reorderList(head *ListNode) {
	var xs []int
	cur := head
	for cur != nil {
		xs = append(xs, cur.Val)
		cur = cur.Next
	}

	p := 0
	q := len(xs) - 1
	cur = head
	for p <= q {
		if cur != nil {
			cur.Val = xs[p]
			cur = cur.Next
		}
		if cur != nil {
			cur.Val = xs[q]
			cur = cur.Next
		}
		p++
		q--
	}
}
