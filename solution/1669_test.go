package solution

import (
	"fmt"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	list1 := []int{0, 1, 2, 3, 4, 5}
	a := 3
	b := 4
	list2 := []int{1000000, 1000001, 1000002}
	list3 := mergeInBetween(NewListNode(list1), a, b, NewListNode(list2))
	fmt.Println(list3)

	list1 = []int{0, 1, 2, 3, 4, 5, 6}
	a = 2
	b = 5
	list2 = []int{1000000, 1000001, 1000002, 1000003, 1000004}
	list3 = mergeInBetween(NewListNode(list1), a, b, NewListNode(list2))
	fmt.Println(list3)
}
