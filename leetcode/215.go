// 数组中第K个最大元素
package leetcode

func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int)int{
	if start >= stop{
			return -1
	}
	pivot := nums[start]
	l, r := start, stop
	for l < r{
			for l < r && nums[r] >= pivot{
					r--
			}
			nums[l] = nums[r]
			for l < r && nums[l] < pivot{
					l++
			}
			nums[r] = nums[l]
	}
	// 循环结束，l与r相等
// 确定基准元素pivot在数组中的位置
	nums[l] = pivot
	return l
}

func TopKSplit(nums []int, k, start, stop int){
	if start < stop{
			index := Parition(nums, start, stop)
			if index == k{
					return
			}else if index < k{
					TopKSplit(nums, k, index+1, stop)
			}else{
					TopKSplit(nums, k, start, index-1)
			}
	}
}