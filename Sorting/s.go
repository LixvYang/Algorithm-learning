func mergesort(arr[]int) {
	len := len(arr)

	if len <= 1 {
		return
	}

	mergeSort(arr,0,len-1)
}

func mergeSort(arr[]int,start,end int) {
	if start >= end {
		return 
	}

	mid := (start + end) /2
	mergeSort(arr,start,mid)
	mergeSort(arr,mid+1,end)

	merge(arr,start,mid,end)
}

func merge(arr[]int,start,mid,end int) {
	tmpArr := make([]int,end-start-1)

	i := start
	j := mid +1
	k := 0

	for ;i<=mid && j <=end;k++{
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		}else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ;i<=mid;i++ {
		tmpArr[k] = arr[i]
		k++
	}

	for ;j<=end;j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start,end+1],tmpArr)
}

// func QuickSort(arr []int) {
// 	separateSort(arr,0,len(arr)-1)
// }

// func separateSort(arr []int,start,end int) {
// 	if start >= end {
// 		return
// 	}

// 	i := partition(arr,start,end)
// 	separateSort(arr,start,i-1)
// 	separateSort(arr,i+1,end)
// }

// func partition(arr []int,start,end int) int {
// 	pivot := arr[end]

// 	var i = start
// 	for j := start;j<end;j++ {
// 		if arr[j] < pivot {
// 			if !(i == j) {
// 				arr[i], arr[j] = arr[j], arr[i]
// 			}
// 			i++
// 		}
// 	}

// 	arr[i],arr[end] = arr[end],arr[i]

// 	return i
// }