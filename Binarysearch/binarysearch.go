package main

import "fmt"

func BinarySearch(a []int, v int) bool {
	n := len(a)

	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high - low)/2

		if a[mid] == v {
			return true
		}else if a[mid] > v {
			high = mid - 1
		}else {
			low  = mid + 1
		}
	}
	if low == len(a) || a[low] != v {
		return false
	}

	return true
}

func main() {
	simple := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(BinarySearch(simple,20))
}