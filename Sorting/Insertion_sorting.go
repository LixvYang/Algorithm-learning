// package sorting

package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 1}
	sort(sample)
	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	sort(sample)
}

// func insertionSort(arr []int) {
//     len := len(arr)
//     for i := 1; i < len; i++ {
//         for j := 0; j < i; j++ {
//             if arr[j] > arr[i] {
//                 arr[j], arr[i] = arr[i], arr[j]
//             }
//         }
//     }
//     fmt.Println("After Sorting")
//     for _, val := range arr {
//         fmt.Println(val)
//     }
// }

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	fmt.Println(arr)
}
