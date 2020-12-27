package solution

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: make([]int, 0), min: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if x <= this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	x := this.data[len(this.data)-1]
	return x
}

func (this *MinStack) GetMin() int {
	x := this.min[len(this.min)-1]
	return x
}
