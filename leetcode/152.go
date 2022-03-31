package leetcode

import "math"

func maxProduct(nums []int) int {
	max := math.IntMin32
	imax := 1
	imin := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := imax
			imax = imin
			imin = temp
		}
		imax = da(nums[i], imax*nums[i])
		imin = xiao(nums[i], imin*nums[i])
		max = da(max, imax)
	}
	return max
}

func da(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func xiao(a, b int) int {
	if a < b {
		return a
	}
	return b
}