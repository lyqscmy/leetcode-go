package solution

func nextLargerNodes(head *ListNode) []int {
	type tuple struct {
		idx int
		val int
	}
	stack := make([]tuple, 0) // 自底向上（index从小到大，val从大到小）还没有next larger的element
	res := make([]int, 0)
	cur := head
	idx := 0
	for cur != nil {
		for len(stack) != 0 && stack[len(stack)-1].val < cur.Val {
			t := stack[len(stack)-1]
			res[t.idx] = cur.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tuple{idx, cur.Val})
		res = append(res, 0)
		cur = cur.Next
		idx++
	}
	return res
}
