package main

import "fmt"

func main() {
	arr := []int{2, 4, 2, 3, 7, 1, 1, 0, 0, 5, 6, 9, 8, 5, 7, 4, 0, 9}
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {
	result := make([]int, len(arr))
	count := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	
	for i, j := 0, 0; i < len(count); i++ {
		for count[i] > 0 {
			result[j] = i 
			j++
			count[i]--
		}
	}
}
