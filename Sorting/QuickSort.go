package main

import "fmt"

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) []int {
	separateSort(arr, 0, len(arr)-1)
	return arr
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	left := start
	right := end-1
	for left < right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for left <= right && arr[right] >= pivot {
			right--
		}
		if (left < right) {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	left, right = right, left
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if arr[i] != arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[end], arr[i] = arr[i], arr[end]
	return i
}
