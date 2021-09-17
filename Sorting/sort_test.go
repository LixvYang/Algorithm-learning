package sorting

import (
	"testing"

	"fmt"
)

func TestBubbleSort(t *testing.T)  {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	BubbleSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestInsertSort(t *testing.T)  {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	InsertSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestSelectSort(t *testing.T)  {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	SelectSort(arr,len(arr))
	fmt.Println("排序后：",arr)
}

func TestMergeSort(t *testing.T)  {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	MergeSort(arr)
	fmt.Println("排序后：",arr)
}

func TestQuickSort(t *testing.T)  {
	arr := []int{1,5,9,6,3,7,5,10}
	fmt.Println("排序前：",arr)
	QuickSort(arr)
	fmt.Println("排序后：",arr)
}