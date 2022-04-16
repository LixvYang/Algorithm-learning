// 滑动窗口的最大值
package leetcode
// 自己建造一个单调队列
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
			queue: []int{},
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
			m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.queue[0] {
			m.queue = m.queue[1:]
	}
}


func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	res := make([]int, 0)
	length := len(nums)
	for i := 0; i < k; i++ {
			queue.Push(nums[i])
	} 
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
			queue.Pop(nums[i-k])
			queue.Push(nums[i])
			res = append(res, queue.Front())
	}
	return res
}