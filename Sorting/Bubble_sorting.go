// 冒泡排序
package main

// import "fmt"

// func main()  {
// 	sample := []int{3, 5, 4, 2, 1}
// 	bubbleSort(sample)
// 	fmt.Println(sample)
//   sample = []int{5, 2, 1, 7, 8, -1, -3}
//   bubbleSort(sample)
// 	fmt.Println(sample)
// }

// func bubbleSort(arr []int)  {
// 	len := len(arr)
// 	for i := 0;i < len - 1; i++ {
// 		for j := 0;j<len-i-1;j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j],arr[j+1] = arr[j+1],arr[j]
// 			}
// 		}
// 	}
// 	fmt.Println("\nAfter Bubblee Sorting")
// 	for _,val := range arr {
// 		fmt.Println(val)
// 	}
// }

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// fmt.Println("\nAfter Bubblee Sorting")
	// for _,val := range arr {
	// 	fmt.Println(val)
	// }
	return arr
}