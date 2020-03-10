package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFindLeftMostBottomNode(t *testing.T) {
	root := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	got := root.findLeftMostBottomNode()
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestHtod(t *testing.T) {
	h := []int{1, 2, 3, 4}
	want := []int{4, 2, 3, 1}
	got := htod(h)
	if Compare(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// func TestDtoh(t *testing.T) {
// 	want := []int{1, 2, 3, 4}
// 	d := []int{4, 2, 3, 1}
// 	got := dtoh(d)
// 	if Compare(got, want) {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

func contains(xs []int, x int) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func TestParseInt(t *testing.T) {
	var tests = []struct {
		in  string
		out int32
	}{
		{"0", 0},
		{"1", 1},
		{strconv.Itoa(1<<32 - 1), -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			act := ParseInt(tt.in)
			if act != tt.out {
				t.Errorf("got %v, want %v", act, tt.out)
			}
		})
	}
}

func TestAddList(t *testing.T) {
	l1 := &ListNode{3, nil}
	l1 = &ListNode{4, l1}
	l1 = &ListNode{2, l1}
	// fmt.Println(l1)

	l2 := &ListNode{4, nil}
	l2 = &ListNode{6, l2}
	l2 = &ListNode{5, l2}
	// fmt.Println(l2)

	exp := &ListNode{8, nil}
	exp = &ListNode{0, exp}
	exp = &ListNode{7, exp}
	// fmt.Println(exp)

	act := addList(l1, l2)
	fmt.Println(act)
	if !act.equal(exp) {
		t.Errorf("got %v, want %v", act, exp)
	}
}
