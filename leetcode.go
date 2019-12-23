package leetcode

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
