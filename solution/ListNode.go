package solution

import (
	"strconv"
	"strings"
)

//ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var b strings.Builder
	for l != nil {
		b.WriteString(strconv.Itoa(l.Val))
		b.WriteString("->")
		l = l.Next
	}
	return b.String()[0 : b.Len()-2]
}

func (l *ListNode) equal(other *ListNode) bool {
	for l != nil && other != nil {
		if l.Val != other.Val {
			return false
		}
		l = l.Next
		other = other.Next
	}

	if l != nil || other != nil {
		return false
	}

	return true
}

func NewListNode(xs []int) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for _, x := range xs {
		cur.Next = &ListNode{x, nil}
		cur = cur.Next
	}
	return dummy.Next
}

func nth(l *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		l = l.Next
	}
	return l
}

func length(l *ListNode) int {
	n := 0
	for l != nil {
		n++
		l = l.Next
	}
	return n
}
