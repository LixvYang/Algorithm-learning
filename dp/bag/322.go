package bag

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0]=0
	for j:=1;j<= amount;j++ {
			dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i:=0;i<len(coins);i++ {
			for j:=coins[i];j<=amount;j++ {
					if dp[j-coins[i]] != math.MaxInt32 {
			// 推导公式
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
			//fmt.Println(dp,j,i)
		}
			}
	}
// 没找到能装满背包的, 就返回-1
if dp[amount] == math.MaxInt32 {
	return -1
}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
			return b
	}
	return a
}