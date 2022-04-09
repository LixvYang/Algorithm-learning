// 完全平方数
package leetcode
// 完全背包问题

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)

    dp[0] = 0
    for i := 1; i <= n; i++ {
        dp[i] = math.MaxInt32
    }

    for i := 1; i <= n; i++ {
        for j := i*i; j <= n; j++ {
            dp[j] = min(dp[j], dp[j-i*i]+1)
        }
    }
    fmt.Println(dp)

    if dp[n] == math.MaxInt32 {
        return -1
    }
    return dp[n]
}	