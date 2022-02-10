// Package leetcode provides 
package leetcode

func canJump(nums []int) bool {
	var cover int 
	for i:=0;i<=cover;i++ {
		cover = max(cover, nums[i]+i)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a,b int) int {
	if a > b {
		return a 
	}
	return b
}