package solution

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	k := 0
	for i, v := range inorder {
		if v == root.Val {
			k = i
			break
		}
	}
	leftInorder := inorder[:k]
	rightInorder := inorder[k+1:]

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[len(preorder)-len(rightInorder):]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
