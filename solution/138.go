package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	node := head
	for node != nil {
		newNode := &Node{Val: node.Val, Next: node.Next, Random: node.Random}
		m[node] = newNode
		node = node.Next
	}

	for _, n := range m {
		n.Next = m[n.Next]
		n.Random = m[n.Random]
	}

	return m[head]
}
