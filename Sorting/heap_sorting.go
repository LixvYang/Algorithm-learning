package main

import "fmt"

// n 数组长度 i待维持下标
func heapify(arr []int, len, i int) {
	largest := i
	lson := i*2 + 1;
	rson := i*2 + 2;

	if lson < len && arr[largest] < arr[lson] {
		largest = lson
	}

	if rson < len && arr[largest] < arr[rson] {
		largest = rson
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, len, largest)
	}
}

func heap_sort(arr []int, n int) {
	// 建堆
	for i := n/2-1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 排序
	for i := n-1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}

}

func main() {
	// arr := []int{14, 7, 10, 3, 4, 9, 8, 2, 1, 16}
	arr := []int{2, 3, 8, 1, 4, 9, 10, 7, 16, 14}
	// arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap_sort(arr, len(arr))
	fmt.Println(arr)
}