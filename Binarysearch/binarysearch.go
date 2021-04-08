package main

import "fmt"

func binarysearch(a []int,v int) int {
	n := len(a)

	if n == 0 {
		return -1
	}

	low := 0
	high := len-1
	for low <= high {
		mid := (low + high) /2
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}