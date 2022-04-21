package bag

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum % 2 != 0 {
		return false
	}
	target := sum/2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}