// 旋转数组中的最小数
package offer

func minArray(numbers []int) int {
	low := 0 
	high := len(numbers)-1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}