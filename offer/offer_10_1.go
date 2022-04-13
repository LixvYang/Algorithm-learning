// 斐波那契数列
package offer

func fib(n int) int {
	if n < 2 {
			return n
	}
	mod := 1000000007
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
			dp[i] = (dp[i-1]+dp[i-2])%mod
	}
	return dp[n]
}