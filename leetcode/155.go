// 最小栈
package leetcode

type MinStack struct {
	arr []int
	min_arr []int
}


func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.arr = append(this.arr, val)
	n := len(this.min_arr)
	if n == 0 || this.min_arr[n-1] > val {
			this.min_arr = append(this.min_arr, val)
	} else {
			this.min_arr = append(this.min_arr, this.min_arr[n-1])
	}
}


func (this *MinStack) Pop()  {
	this.arr = this.arr[:len(this.arr)-1]
	this.min_arr = this.min_arr[:len(this.min_arr)-1]
}


func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
	return this.min_arr[len(this.min_arr)-1]
}