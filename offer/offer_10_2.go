// 青蛙跳台阶问题
package offer


func numWays(n int) int {
	if n == 0 {
			return 1
	}
	if n == 1 {
			return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
			dp[i] = (dp[i-1]+dp[i-2])%1000000007
	}
	return dp[n]
}