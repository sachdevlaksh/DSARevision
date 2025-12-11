package LC155


type node struct {
	min int
	val int
}
type MinStack struct {
	stack []node
}

func Constructor() MinStack {
	return MinStack{stack: make([]node, 0)}
}

func (this *MinStack) Push(val int) {

	if len(this.stack) > 0 {
		currMin := this.stack[len(this.stack)-1].min
		newNode := node{val: val}
		if currMin > val {
			newNode.min = val
		} else {
			newNode.min = currMin
		}
		this.stack = append(this.stack, newNode)
	} else {
		this.stack = append(this.stack, node{val: val, min: val})
	}

}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1].val
	} else {
		return 0
	}
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1].min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
