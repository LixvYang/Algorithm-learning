package main

import (
	"errors"
	"fmt"
)

// 小顶堆
type Heap struct {
	nums []int
}

func NewHeap() *Heap {
	return &Heap{
		nums: make([]int, 1),
	}
}

func (h *Heap) Push(val int) {
	h.nums = append(h.nums, val)
	pos := len(h.nums) - 1

	for pos > 1 && h.nums[pos/2] > h.nums[pos] {
		h.nums[pos/2], h.nums[pos] = h.nums[pos], h.nums[pos/2]
		pos = pos / 2
	}
}

func (h *Heap) Pop() (int, error) {
	if len(h.nums) <= 1 {
		return 0, errors.New("Empty")
	}

	val := h.nums[1]
	h.nums[1], h.nums[len(h.nums)-1] = h.nums[len(h.nums)-1], h.nums[1]
	h.nums = h.nums[:len(h.nums)-1]
	if len(h.nums) > 1 {
		h.adjustDown(1)
	}
	return val, nil
}

func (this *Heap) adjustDown(pos int) {
	length := len(this.nums) - 1
	//此处规定下标从1开始，因此下标为0的位置可以用来暂时保存值
	this.nums[0] = this.nums[pos]
	//for循环为找子节点过程
	for i := 2 * pos; i <= length; i *= 2 {
		//找到更小的子节点
		if i < length && this.nums[i] > this.nums[i+1] {
			i++
		}
		//如果存在子节点小于需要调整的结点，则将子节点移动到父节点位置
		if this.nums[0] > this.nums[i] {
			this.nums[pos] = this.nums[i]
			pos = i
		}
	}
	//将结点放置在正确的位置
	this.nums[pos] = this.nums[0]
}

//测试
func main() {
	heap := NewHeap()
	heap.Push(1)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	fmt.Println(heap.nums)
	pop, _ := heap.Pop()
	fmt.Println(pop) //1
	pop, _ = heap.Pop()
	fmt.Println(pop) //2
	pop, _ = heap.Pop()
	fmt.Println(pop) //3
	pop, _ = heap.Pop()
	fmt.Println(pop) //4
}
