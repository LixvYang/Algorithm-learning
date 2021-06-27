package stack

import (
	"fmt"
)

type ArrayStack struct {
	//data
	data []int
	//top
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:-1,
	}
}

func (this *ArrayStack) isEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v interface{})  {
	if this.top < 0{
		this.top = 0
	} else {
		this.top += 1
	}

	if this.top > len(this.data) - 1{
		this.data = append(this.data,v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.isEmpty() {
		return nil
	}

	v := this.data[this.top]
	this.top-=1
	return v
}

func (this *ArrayStack) Top() interface{} {
	if this.isEmpty() {
		return nil
	}

	return this.data[this.top]
}

func (this *ArrayStack) Flush()  {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.isEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := this.top;i>=0;i-- {
			fmt.Println(this.data[i])
		}
	}
}