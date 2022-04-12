// 最短无序连续子数组
package leetcode

func findUnsortedSubarray(nums []int) int {
	// 返回最短无序子数组
	length := len(nums)
	min := nums[length-1] // 15
	max := nums[0] // 2
	// 无序数组的左右边界 
	begin, end := 0, -1
	// 遍历
	for i:=0;i<length;i++ {
		// 一个更新min或max，一个更新左右边界
		// 从左往右维持最大值，寻找右边界end
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
		// 从右至左寻找最小值，寻找左边界begin
		if nums[length-i-1] > min {
			begin = length-i-1
		} else {
			min = nums[length-1-i]
		}
	}
	return end-begin+1
}

