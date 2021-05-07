package solution

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head = sortList2(head)
	fmt.Println(head)
}
