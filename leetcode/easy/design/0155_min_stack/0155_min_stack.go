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

// 每次放入两个元素：{min, v}
// v 是最小时放入 {v, v}
// v 不是最小时放入 {min, v}
func (this *MinStack) Push(v int) {
	min := v
	if len(this.stack) > 0 && this.GetMin() < v {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min, v})
}

// 删除数组最后一个元素
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

// 返回数组中最后一个元素的 value 值
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

// 返回数组中最后一个元素的 min 值
func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
