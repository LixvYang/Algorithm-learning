package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]

	mx := nums[0]

	for i:=1;i<n;i++ {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
			mx = max(dp[i], mx)
	}
	return mx
}

func max(a, b int) int {
	if a < b {
			return b 
	}
	return a
}