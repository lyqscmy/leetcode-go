package leetcode

import (
	"strconv"
	"strings"
)

func Search(xs []int, x int) int {
	i, j := 0, len(xs)
	for i < j {
		// i<=h<j
		h := int(uint(i+j) >> 1)
		if xs[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func findMin(nums []int) int {
	i, j := 0, len(nums)
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[i] < nums[h] {
			i = h + 1
		} else {
			j = h
		}

	}
	return i
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) equal(other *TreeNode) bool {
	return false
}

func (root *TreeNode) findLeftMostBottomNode() int {
	res := []int{-1}
	res = root.findLeftMostBottomNodeImpl(1, res)
	return res[len(res)-1]
}

func (root *TreeNode) findLeftMostBottomNodeImpl(index int, res []int) []int {
	if root == nil {
		return res
	}

	if index > len(res) {
		res = append(res, root.Val)
	}

	res = root.Left.findLeftMostBottomNodeImpl(index+1, res)
	res = root.Right.findLeftMostBottomNodeImpl(index+1, res)
	return res
}

type Deque struct {
	items  []int
	head   int
	tail   int
	length int
}

func NewDeque(n int) *Deque {
	return &Deque{make([]int, n+1), 0, 0, 0}
}

func (d *Deque) AddFirst(e int) {
	d.length++
	d.head = dec(d.head, len(d.items))
	d.items[d.head] = e
}

func dec(i int, modulus int) int {
	if i == 0 {
		i = modulus - 1
	} else {
		i--
	}
	return i
}
func (d *Deque) AddLast(e int) {}

// func (d *Deque) RemoveFirst() int {}
func (d *Deque) RemoveLast() int {
	d.length--
	d.tail = dec(d.tail, len(d.items))
	return d.items[d.tail]
}
func (d *Deque) Len() int { return d.length }
func (d *Deque) Items() []int {
	items := make([]int, 0, d.length)
	head := d.head
	for head != d.tail {
		items = append(items, d.items[head])
		head = dec(head, len(d.items))
	}
	return items
}

// left to right = top to bottom
func dtoh(d []int) []int {
	n := len(d)
	h := NewDeque(n)
	for i := 0; i < n; i++ {
		// step 2
		if h.Len() != 0 {
			h.AddFirst(h.RemoveLast())
		}

		// step 1
		h.AddFirst(d[i])
	}
	return h.Items()
}

// left to right = top to bottom
func htod(h []int) []int {
	n := len(h)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		// step 1
		d[i] = h[0]
		h = h[1:]

		// step 2
		if len(h) != 0 {
			x := h[0]
			h = h[1:]
			h = append(h, x)
		}
	}
	return d
}

func Compare(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findPeakElement(nums []int) int {
	N := len(nums)
	for i := 0; i < N-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return N - 1
}

func ParseInt(s string) int32 {
	var n int32 = 0
	var cutoff int32 = (1<<32)/10 + 1
	for i := 0; i < len(s); i++ {
		if n >= cutoff {
			return -1
		}
		n *= 10
		n1 := n + int32(s[i]-'0')
		if n1 < n {
			return -1
		}
		n = n1
	}
	return n
}

func Search(xs []int, x int) int {
	i, j := 0, len(xs)
	for i < j {
		// i<= h < j
		h := int(uint(i+j) >> 1)
		if x > xs[h] {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) String() string {
	var b strings.Builder
	for this != nil {
		b.WriteString(strconv.Itoa(this.Val))
		b.WriteString("->")
		this = this.Next
	}
	return b.String()[0 : b.Len()-2]
}

func (this *ListNode) equal(other *ListNode) bool {
	for this != nil && other != nil {
		if this.Val != other.Val {
			return false
		}
		this = this.Next
		other = other.Next
	}

	if this != nil || other != nil {
		return false
	}

	return true
}

func addList(l1 *ListNode, l2 *ListNode) *ListNode {
	dumy := &ListNode{}
	node := dumy
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		node = node.Next
		carry /= 10
	}
	if carry != 0 {
		node.Next = &ListNode{carry, nil}
	}
	return dumy.Next
}
