package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
			return 0
	}

	dp := make([][]int, n)
	status := make([]int, n * 4)
	for i := range dp {
			dp[i] = status[:4]
			status = status[4:]
	}
	dp[0][0] = -prices[0]

	for i:=1;i<n;i++ {
			dp[i][0] = max(dp[i - 1][0], max(dp[i - 1][1] - prices[i], dp[i - 1][3] - prices[i]))
			dp[i][1] = max(dp[i - 1][1], dp[i - 1][3])
			dp[i][2] = dp[i - 1][0] + prices[i]
			dp[i][3] = dp[i - 1][2]
	}
	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

func max(a, b int) int {
	if a<b{
			return b
	} else {
			return a
	}
}