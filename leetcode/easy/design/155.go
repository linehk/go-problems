package design

type MinStack struct {
	stack []item
}

type item struct {
	min   int
	value int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(v int) {
	min := v
	if len(this.stack) > 0 && this.GetMin() < v {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min, v})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
