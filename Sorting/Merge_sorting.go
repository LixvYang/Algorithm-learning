package sorting
// package main

import (
	"fmt"
	"time"
)
//主函数程序
func main()  {
	//赋值
	sample := []int{3, 4, 5, 2, 1,6,7,8}
	//调用函数
	mergesort(sample)
}

func mergesort(arr []int)  {
	starttime := time.Now()
	len := len(arr)

	if len <=1 {
		return
	}

	mergeSort(arr,0,len-1)
	fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
	elapsed := time.Since(starttime)
  fmt.Println("All time: ", elapsed)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmp)
}

// func sort(arr []int) {

// }

// func mergeSort(arr []int, start, end int) {
// 	if start >= end {
// 		return
// 	}

// 	mid := start+(end-start)/2
// 	mergeSort(arr, start, mid)
// 	mergeSort(arr, mid+1, end)
// 	merge(arr, start, mid, end)
// }

// func merge(arr []int, start, mid, end int) {
// 	temp := make([]int, end-start+1)

// 	i, j := start, mid+1
// 	k := 0

// 	for i <= mid && j <= end {
// 		if arr[i] <= arr[j] {
// 			temp[k] = arr[i]
// 			i++
// 			k++
// 		}  else {
// 			temp[k] = arr[j]
// 			j++
// 			k++
// 		}
// 	}

// 	for i <= mid {
// 		temp[k] = arr[i]
// 		i++
// 		k++
// 	}
// 	for j < len(arr) {
// 		temp[k] = arr[j]
// 		j++
// 		k++
// 	}
// 	for k, v := range temp {
// 		arr[k] = v
// 	}
// }