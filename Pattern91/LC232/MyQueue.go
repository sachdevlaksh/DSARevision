/*
LeetCode Problem #232: Implement Queue using Stacks
Difficulty: Easy

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue.
*/

package LC232

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{
		make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append([]int{x}, this.stack...)
}

func (this *MyQueue) Pop() int {
	if len(this.stack) > 0 {
		c := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return c
	}
	return 0
}

func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {

	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
