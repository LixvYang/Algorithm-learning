package leetcode

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i:=0;i<len(prices);i++ {
			dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]

	for i:=1;i<len(prices);i++ {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
			dp[i][1] = max(dp[i-1][0], dp[i][0]+prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a < b {
			return b
	}
	return a
}