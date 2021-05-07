package solution

import (
	"sort"
)

func sortList(head *ListNode) *ListNode {
	ls := make([]*ListNode, 0)
	for head != nil {
		ls = append(ls, head)
		tmp := head
		head = head.Next
		tmp.Next = nil
	}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i].Val < ls[j].Val
	})
	dummy := &ListNode{}
	node := dummy
	for _, l := range ls {
		node.Next = l
		node = node.Next
	}
	return dummy.Next
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		fast = fast.Next
		slow = slow.Next
	}
	q := sortList2(slow.Next)
	slow.Next = nil
	p := sortList2(head)

	dummy = &ListNode{}
	node := dummy
	for q != nil || p != nil {
		switch {
		case q != nil && p != nil:
			if q.Val < p.Val {
				node.Next = &ListNode{q.Val, nil}
				q = q.Next
			} else {
				node.Next = &ListNode{p.Val, nil}
				p = p.Next
			}
		case q != nil && p == nil:
			node.Next = &ListNode{q.Val, nil}
			q = q.Next
		case p != nil && q == nil:
			node.Next = &ListNode{p.Val, nil}
			p = p.Next
		}
		node = node.Next
	}
	return dummy.Next
}
