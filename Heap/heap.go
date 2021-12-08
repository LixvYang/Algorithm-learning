package heap

type Heap struct {
	a	[]int	//数组 从1 开始存储
	n int // 堆可以存储的最大数据的个数
	count int // 堆已经存储的数据个数
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity + 1)
	heap.count = 0
	return heap
}

func (heap *Heap) insert(data int) {
	if heap.count == heap.n {
		return 
	}

	heap.count++
	heap.a[heap.count] = data

	i := heap.count
	parent := i/2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i/2
	}	
}

func swap(a []int, i,j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}


func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}

	swap(heap.a, 1, heap.count)
	heap.count--

	//heapify from up to down
	heapify(heap.a, heap.count)
}

func (heap *Heap) heapify(a []int, count int)  {
	for i := 1; i <= count/2 {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 < count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2+1
		}

		swap(a, i, maxIndex)

		i = maxIndex
	}
}
