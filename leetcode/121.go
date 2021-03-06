// 买卖股票的最佳时机
package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -price[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] + price[0])
	}
	return dp[len(price)-1][1]
}   

func max(a, b int) int {
	if a < b {
			return b
	}
	return a
}