// Package leetcode provides 
package leetcode

type MyQueue struct {
	stack	[]int
	back	[]int
}

func Constructor() MyQueue {
	return MyQueue{
			stack: make([]int, 0),
			back: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) int {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}

	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.back[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}

	if len(this.back) == 0 {
		return 0
	}

	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

func (this *MyQueue) Peek() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
			return 0
	}
	val := this.back[len(this.back)-1]
	return val
}

func (this *MyQueue) IsEmpty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}