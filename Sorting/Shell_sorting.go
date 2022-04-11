// 希尔排序
package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0}
	shellSort(arr)
	fmt.Println(arr)
}

func shellSort(arr []int) {
	for gap := 4; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >gap-1;j-=gap {
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}
			}
		}
	}
	
}