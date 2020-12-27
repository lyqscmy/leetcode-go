package solution

import (
	"encoding/json"
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
	str, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(str)
}
