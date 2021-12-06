package binarysearch

import (
	"fmt"
	"time"
)

func BinarySearch(a []int, v int) bool {
	
	n := len(a)

	low := 0
	high := n - 1
	starttime := time.Now()
	for low <= high {
		mid := low + (high-low)/2
		
		if a[mid] == v {
			elapsed := time.Since(starttime)
    		fmt.Println("All time: ", elapsed)
			return true
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	elapsed := time.Since(starttime)
    fmt.Println("All time: ", elapsed)

	return false
}

func main() {
	simple := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(BinarySearch(simple, 45))
}