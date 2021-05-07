package solution

func numComponents(head *ListNode, G []int) int {
	set := make(map[int]bool, len(G))
	for _, val := range G {
		set[val] = true
	}
	res := 0
	for head != nil {
		if set[head.Val] && (head.Next == nil || !set[head.Next.Val]) {
			res++
		}
		head = head.Next
	}
	return res
}
