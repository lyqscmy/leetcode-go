package solution

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	got := sortedArrayToBST(nums)
	want := &TreeNode{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s , want: %s", got, want)
	}

	nums = []int{}
	got = sortedArrayToBST(nums)
	want = &TreeNode{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s , want: %s", got, want)
	}
}
