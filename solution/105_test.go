package solution

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

func TestBuildTree2(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}
