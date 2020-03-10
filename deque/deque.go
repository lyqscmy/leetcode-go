package deque

type Deque struct {
	items []int
	head  int
	tail  int
	size  int
}

// NewDeque
func NewDeque(capcity int) *Deque {

}

func (d *Deque) Size() int {
	if d == nil {
		return 0
	}
	return d.size
}

// Push put val back of the deque, grow when needed
func (d *Deque) Push(val int) {
}

// Pop remove the first val of the deque and return it
func (d *Deque) Pop() int {
}

// Peek return first val of the deque
func (d *Deque) Peek() int {
}

func inc(x int, modulus int) int {
}

func dec(x int, modulus int) int {
}

func grow(x int, modulus int) int {
}
