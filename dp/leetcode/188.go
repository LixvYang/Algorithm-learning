package leetcode

func maxProfit(k int, prices []int) int {
	if len(prices)==0{
			return 0
	}
	dp:=make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
			dp[i]=make([]int,2*k+1)
	}
	for i:=1;i<len(dp[0]);i++{
			if i%2!=0{
					dp[0][i]=-prices[0]
			}
	}

	for i:=1;i<len(prices);i++ {
			dp[i][0] = dp[i-1][0]
			for j:=1;j<len(dp[0]);j++ {
					if j%2!=0 {
							dp[i][j] = max(dp[i-1][j],dp[i-1][j-1]-prices[i])
					} else {
							dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
					}
			}
	}
	return dp[len(prices)-1][2*k]
}

func max(a,b int) int {
	if a> b {
			return a
	}
	return b
}