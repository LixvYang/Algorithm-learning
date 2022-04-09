// 最佳买卖股票时机含冷冻期
package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
  if n < 2 {
      return 0
  }

	dp := make([][]int, n)
	for i := dp {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0]
	// 第一种状态： 持有股票
	// 第二种状态： 卖出股票状态 度过了冷冻期，一直没操作，今天保持卖出股票状态
	//	第三种状态： 卖出股票状态 今天卖出股票
	//	第四种状态： 今天为冷冻期
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0]+price[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}