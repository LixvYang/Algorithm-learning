// 最长递增子序列
package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	result := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}

			if result < dp[i] {
				result = dp[i]
			}
		}
	}
	return result
}