package solution

import (
	"fmt"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	level := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root, p, q := constructTree(level, 5, 4)
	fmt.Println(root)
	lca := lowestCommonAncestor(root, p, q)
	fmt.Println(lca)
}

func constructTree(level []int, p int, q int) (*TreeNode, *TreeNode, *TreeNode) {
	type tuple struct {
		index int
		root  *TreeNode
	}
	root := &TreeNode{Val: level[0]}
	queue := make([]*tuple, 0)
	queue = append(queue, &tuple{0, root})
	var pn, qn *TreeNode
	if root.Val == p {
		pn = root
	}
	if root.Val == q {
		qn = root
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left := node.index*2 + 1
		if left < len(level) && level[left] != -1 {
			node.root.Left = &TreeNode{Val: level[left]}
			queue = append(queue, &tuple{left, node.root.Left})
			if node.root.Left.Val == p {
				pn = node.root.Left
			}
			if node.root.Left.Val == q {
				qn = node.root.Left
			}
		}
		right := node.index*2 + 2
		if right < len(level) && level[right] != -1 {
			node.root.Right = &TreeNode{Val: level[right]}
			queue = append(queue, &tuple{right, node.root.Right})
			if node.root.Right.Val == p {
				pn = node.root.Right
			}
			if node.root.Right.Val == q {
				qn = node.root.Right
			}
		}
	}
	return root, pn, qn
}
