package leetcode

// Deque is
type Deque struct {
	items  []int
	head   int
	tail   int
	length int
}

// NewDeque do
func NewDeque(n int) *Deque {
	return &Deque{make([]int, n+1), 0, 0, 0}
}

// AddFirst do
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

// AddLast do
func (d *Deque) AddLast(e int) {}

// RemoveLast do
func (d *Deque) RemoveLast() int {
	d.length--
	d.tail = dec(d.tail, len(d.items))
	return d.items[d.tail]
}

// Len return
func (d *Deque) Len() int { return d.length }

// Items return
func (d *Deque) Items() []int {
	items := make([]int, 0, d.length)
	head := d.head
	for head != d.tail {
		items = append(items, d.items[head])
		head = dec(head, len(d.items))
	}
	return items
}