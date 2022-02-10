// Package leetcode provides 
package leetcode

func wiggleMaxLength(nums []int) int {
	var count, pre, cur int
	count = 1
	if len(nums)<2 {
			return count
	}
	
	for i := 0; i < len(nums)-1; i++ {
			cur = nums[i+1]-nums[i]
			if (cur>0&&pre<=0)||(cur<0&&pre>=0) {
					pre = cur
					count++
			}
	}
	return count
}
