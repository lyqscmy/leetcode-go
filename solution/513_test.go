package solution

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestFindLeftMostBottomNode(t *testing.T) {
	blob := []byte(`{ "val": 2, "left": { "val": 1 }, "right": { "val": 3 } }`)
	t1 := &TreeNode{}
	err := jsoniter.Unmarshal(blob, t1)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		root *TreeNode
		want int
	}{
		{t1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findLeftMostBottomNode(tt.root); got != tt.want {
				t.Errorf("findLeftMostBottomNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
