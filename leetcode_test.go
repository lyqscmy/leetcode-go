package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestDtoh(t *testing.T) {
	want := []int{1, 2, 3, 4}
	d := []int{4, 2, 3, 1}
	got := dtoh(d)
	if Compare(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindPeakElement(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 1}, []int{2}},
		{[]int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			act := findPeakElement(tt.in)
			if !contains(tt.out, act) {
				t.Errorf("got %v, want %v", act, tt.out)
			}
		})
	}
}

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

func TestSearch(t *testing.T) {
	var tests = []struct {
		xs  []int
		x   int
		exp int
	}{
		{[]int{}, 1, 0},
		{[]int{1}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			act := Search(tt.xs, tt.x)
			if act != tt.exp {
				t.Errorf("got %v, want %v", act, tt.exp)
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

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, []int{1, 2, 2, 3, 5, 6}},
		{"", args{nums1: []int{4, 5, 6, 0, 0, 0}, m: 3, nums2: []int{1, 2, 3}, n: 3}, []int{1, 2, 3, 4, 5, 6}},
		{"", args{nums1: []int{0, 0, 0, 0, 0}, m: 0, nums2: []int{1, 2, 3, 4, 5}, n: 5}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !Compare(tt.args.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
