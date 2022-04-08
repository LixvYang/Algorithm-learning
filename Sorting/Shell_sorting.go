// 希尔排序
package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0}
	shellSort(arr)
}

func shellSort(arr []int) {
	length := len(arr)

	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
 	}
	fmt.Println(gap)

	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i] 
			j := i-gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j-=gap
			}
			arr[j+gap] = temp
		}
		gap /= 3
	}
	fmt.Println(arr)
}