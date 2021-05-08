package main

import (
	"fmt"
)

func main() {
	

}

func QuickSort(arr []int)  {
	len := len(arr)

	separateSort(arr,0,len-1)

	fmt.Println("After Sorting")

	for _,value := range arr {
		fmt.Println(value)
	}
}

func sparateSort(arr []int,start,end int)  {
	if start >= end {
		return
	}

	i := partition(arr,start,end)

	separateSort(arr,start,i-1)
	separateSort(arr,i+1,end)
}

func paratition(arr []int,start,end int) int {
	pivot := arr[end]

	var i = start
	for j := start;j<end;j++ {
		if arr[j] < arr[end] {
			if !(i == j) {
				arr[i],arr[j] = arr[j],arr[i]
			}
			i++
		}

		arr[i],arr[end] = arr[end],arr[i]
		return i
	}
}