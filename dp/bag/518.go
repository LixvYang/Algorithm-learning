package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
		fmt.Println(dp)
	}

	for j := 0; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
		fmt.Println(dp)
	}
	return dp[amount]
}

func main() {
	change(5, []int{1, 2, 5})
}
