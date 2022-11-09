package main

import "fmt"

func climbStairs(n int, m int) int {
	//定义
	dp := make([]int, n+1)
	//初始化
	dp[0] = 1
	// 本题物品只有两个1,2
	// m := 2
	// 遍历顺序
	for j := 1; j <= n; j++ { //先遍历背包
		for i := 1; i <= m; i++ { //再遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
			fmt.Println(dp)
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(20, 2))
}
