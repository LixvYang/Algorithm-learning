// 希尔排序
// package main

// import "fmt"

// func main() {
// 	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0}
// 	shellSort(arr)
// 	fmt.Println(arr)
// }

func shellSort(arr []int) {
	h := 1
	for h <= len(arr)/3 {
		h = h*3+1
	}


	for gap := h; gap > 0; gap = (gap-1)/3 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >gap-1;j-=gap {
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}
			}
		}
	}
}
