package leetcode

import (
	"strings"
	"strconv"
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