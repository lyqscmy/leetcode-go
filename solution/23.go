package solution

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	for i := 0; i < len(lists); i++ {
		head := lists[i]
		if head != nil {
			heap.Push(&pq, &Item{value: i, priority: head.Val})
			head = head.Next
			lists[i] = head
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		tail.Next = &ListNode{item.priority, nil}
		tail = tail.Next
		head := lists[item.value]
		if head != nil {
			item.priority = head.Val
			heap.Push(&pq, item)
			head = head.Next
			lists[item.value] = head
		}
	}
	return dummy.Next
}
