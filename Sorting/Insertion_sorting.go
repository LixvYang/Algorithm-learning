package sorting

// package main

// import "fmt"

// func main() {
//     sample := []int{3, 4, 5, 2, 1}
//     insertionSort(sample)

//     sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
//     insertionSort(sample)
// }

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
	len := len(arr)
	for i := 1; i < len; i++ {
		value := arr[i]
		j := i-1
		for ;j > 0;j-- {
			if value < arr[j] {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}
