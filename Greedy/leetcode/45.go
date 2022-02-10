// Package leetcode provides 
package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
			return 0
	}

	var curDistance, nextDistance, ans int
	for i := 0; i < len(nums); i++ {
			nextDistance = max(nextDistance, nums[i]+i)
			if i == curDistance {
					if curDistance != len(nums)-1 {
							ans++
							curDistance = nextDistance
							if nextDistance >= len(nums)-1 {
									break
							}
					} else {
							break
					}
			}
	}
	return ans
}