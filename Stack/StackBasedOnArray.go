package stack

import (
	"fmt"
)
type ArrayStack struct {
	data []interface{}
	top int
}

func  NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data:make([]interface{}, 0,32),
		top:-1,
	}
}

func (a *ArrayStack) isEmpty() bool {
	if a.top < 0 {
		return true
	}
	return false
}

func (a *ArrayStack) Push(v interface{})  {
	if a.top < 0 {
		a.top = 0
	} else {
		a.top+=1
	}

	if a.top > len(a.data)-1 {
		a.data = append(a.data, v)
	}else {
		a.data[a.top] = v
	}
}

func (a *ArrayStack) Pop() interface{} {
	if a.isEmpty() {
		return nil
	}
	v := a.data[a.top]
	a.top--
	return v
}

func (this *ArrayStack) Top() interface{} {
	if this.isEmpty() {
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.isEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}