package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
			return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
			dp[i] = 1
	}

	result := 0
	for i:=0;i<len(nums);i++ {
			for j:=0;j<i;j++ {
					if nums[i]>nums[j]  && dp[j] + 1 > dp[i]{
							dp[i] = max(dp[i], dp[j]+1)
					}
			}
			if dp[i] > result {
					result = dp[i]
			}
	}
	return result
}

func max(a, b int) int {
	if a < b {
			return b
	}
	return a
}