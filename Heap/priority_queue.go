package heap

type Node struct {
	value int
	priority int
}

func adjustHeap(src []Node,start,end int) {
	if start >= end {
		return
	}

	for i := end/2; i >= start; i-- {
		high := i
		if src[high].priority <= src[i*2].priority {
			high = i*2
		}
		if 2*i+1 <= end && src[high].priority < src[2*i+1].priority {
			high = 2*i + 1
		}
		if high == i {
			continue
		}

		src[high], src[i] = src[i], src[high]
	} 
}

type PQueue struct {
	heap []Node

	capacity int
	used int
}

func NewPQueue(capacity int) PQueue {
	return PQueue{
		heap: make([]Node, capacity+1,capacity+1),
		capacity: capacity,
		used:0
	}
}

// 入队
func (q *PQueue) Push(node Node) {
	if q.used > q.capacity {
		return 
	}

	q.used++
	q.hap[q.used] = node

}

func (q *PQueue) Pop() Node {
	if q.used == 0 {
		return Node{-1,-1}
	}

	adjustHeap(q.heap,1,q.used)
	node := q.heap[1]
	q.heap[1] = q.heap[used]
	q.used--
	return node
}

func (q *PQueue) Top() Node {
	if q.used == 0 {
		return Node{-1,-1}
	}

	adjustHeap(q.heap, q.used)
	return q.heap[1]
}