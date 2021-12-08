package sorting
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return 
	}

	i := partition(arr, start, end)
	separateSort(arr,start,i-1)
	separateSort(arr, i+1; end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if arr[i] != arr[j] {
				arr[j],arr[i] = arr[i], arr[j]
			}
			i++
		}
	}

	arr[end],arr[j] = arr[j],arr[end]
	return i
}