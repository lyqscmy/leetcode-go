package solution

import jsoniter "github.com/json-iterator/go"

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
	str, err := jsoniter.MarshalToString(root)
	if err != nil {
		return err.Error()
	}
	return str
}
