package main

import "fmt"
//主函数程序
func main()  {
	//赋值
	sample := []int{3, 4, 5, 2, 1,6,7,8}
	//调用函数
	mergesort(sample)
}

func mergesort(arr []int)  {
	len := len(arr)

	if len <= 1 {
		return
	}

	mergeSort(arr,0,len-1)

	fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}

func mergeSort(arr []int,start,end int) {
	//前提
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr,start,mid)
	mergeSort(arr,mid+1,end)

	merge(arr,start,mid,end)
}

func merge(arr []int,start,mid,end int)  {
	tmpArr := make([]int,end-start+1)

	i := start
	j := mid +1
	k := 0

	for ;i <= mid && j <= end;k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
 	}
	//剩下的两个循环是有一个i或者j无法被比较
	//于是就如此  让最后一个值赋值到tmpArr[k]中
	for ;i <= mid;i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ;j <= end;j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1],tmpArr)
}