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
	// for len(this.back) != 0 {
	// 	val := this.back[len(this.back)-1]
	// 	this.back = this.back[:len(this.back)-1]
	// 	this.stack = append(this.stack, val)
	// }

	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
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
	// for len(this.stack) != 0 {
	// 	val := this.stack[len(this.stack)-1]
	// 	this.stack = this.stack[:len(this.stack)-1]
	// 	this.back = append(this.back, val)
	// }
	// if len(this.back) == 0 {
	// 		return 0
	// }
	// val := this.back[len(this.back)-1]
	// return val

	res := this.Pop()
	this.back = append(this.back, res)
	return res
}

func (this *MyQueue) IsEmpty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

if len(this.stackB) == 0 {
	// 如果 stackA 里也没有元素，则队列爲空返回 -1
	if len(this.stackA) == 0 {
			return -1
	}
	// 将 stackA 的所有元素转移到 stackB
	for len(this.stackA) > 0 {
			// 获取 stackA 最末尾元素的下标
			index := len(this.stackA) - 1
			// 获取 stackA 最末尾元素的值 value
			value := this.stackA[index]
			// 向 stackB 的末尾追加 value
			this.stackB = append(this.stackB, value)
			// 从 stackA 中裁剪出末尾元素
			this.stackA = this.stackA[:index]
	}
}
