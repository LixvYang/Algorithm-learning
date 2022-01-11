// Package leetcode provides 
package leetcode

type MyStack struct {
	queue []int
}


func Constructor() MyStack {
	return MyStack{
			queue: make([]int, 0),
	}
}


func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
	n := len(this.queue) - 1
	for n != 0 {
			val := this.queue[0]
			this.queue = this.queue[1:]
			this.queue = append(this.queue, val)
			n--
	}

	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}


func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}


func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
