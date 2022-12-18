package leetcode

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' { // ...(
			dp[i] = 0
		} else if s[i-1] == '(' { // ...()
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if dp[i-1] > 0 { //   ..(..))
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' { // .((..))
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2] // (..)((..))
				}
			}

		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
