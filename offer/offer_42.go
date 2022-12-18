package offer

func maxSubArray(nums []int) int {
	// 设动态规划列表dp,dp[i] 代表以元素nums[i]为结尾的连续子数组最大和
	// 以某个数作为结尾，意思就是这个数一定会加上去，那么要看的就是这个数前面的部分要不要加上去。大于零就加，小于零就舍弃。
	// dp := []int{nums[0]}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxValue := dp[0]
	// 转移方程:
	// (1)当 dp[i - 1] > 0 时：执行 dp[i] = dp[i-1] + nums[i]
	// (2)当 dp[i−1]≤0 时：执行 dp[i] = nums[i]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if maxValue < dp[i] {
			maxValue = dp[i]
		}
	}
	return maxValue
}
