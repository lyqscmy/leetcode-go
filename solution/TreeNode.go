package solution

import (
	"strconv"
	"strings"
)

//TreeNode is
type TreeNode struct {
	Val   int       `json:val`
	Left  *TreeNode `json:left`
	Right *TreeNode `json:right`
}

func (root *TreeNode) equal(other *TreeNode) bool {
	return false
}

func (root *TreeNode) String() string {
	var sb strings.Builder
	root.print(&sb, "", "")
	return sb.String()
}

func (root *TreeNode) print(sb *strings.Builder, prefix, childrenPrefix string) {
	sb.WriteString(prefix)
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteRune('\n')
	if root.Left != nil {
		root.Left.print(sb, childrenPrefix+"├── ", childrenPrefix+"│   ")
	}
	if root.Right != nil {
		root.Right.print(sb, childrenPrefix+"└── ", childrenPrefix+"    ")
	}
}
