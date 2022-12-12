package main

import "fmt"

// n 数组长度 i待维持下标
func heapify(arr []int, len, i int) {
	largest := i
	lson := i*2 + 1
	rson := i*2 + 2

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
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 排序
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}

}

func main() {
	// arr := []int{14, 7, 10, 3, 4, 9, 8, 2, 1, 16}
	arr := []int{2, 3, 8, 1, 4, 9, 10, 7, 16, 14}
	// arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(len(arr))
	heap_sort(arr, len(arr))
	fmt.Println(arr)
}

//         1
//			 /  \
//			2    3
//		 / \   / \
//		4  5  6   9    i*2+2

// func h_sort(arr []int, len int) {
// 	// 建堆
// 	for i := len/2 - 1; i >= 0; i-- {
// 		h(arr, len, i)
// 	}
// 	// 堆排序
// 	for i := len - 1; i >= 0; i-- {
// 		arr[i], arr[0] = arr[0], arr[i]
// 		h(arr, len, 0)
// 	}
// }

// func sort(arr []int) {
// 	merge_sort(arr, 0, len(arr)-1)
// }

// func merge_sort(arr []int, start, end int) {
// 	if start == end {
// 		return
// 	}
// 	mid := (start + end) /2
// 	merge_sort(arr, start, mid)
// 	merge_sort(arr, mid+1, end)
// 	merge(arr, start, mid, end)
// }

// func merge(arr []int, start, mid, end int) {
// 	i, j, k := start, mid+1, 0
// 	tmpArr := make([]int, j-i+1)
// 	for i <= mid && j <= end; k++ {
// 		if arr[i] <= arr[j] {
// 			tmpArr[k] = arr[i]
// 			i++
// 		} else {
// 			tmpArr[k] = arr[j]
// 			j++
// 		}
// 	}

// 	for ; i <= mid; i++ {
// 		tmpArr[k] = arr[i]
// 		k++
// 	}

// 	for ; j <= end; i++ {
// 		tmpArr[k] = arr[j]
// 		k++
// 	}
// 	for i := 0; i < len(tmpArr); i++ {
// 		arr[start+i] = tmpArr[i]
// 	}
// }
