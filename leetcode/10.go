// 正则表达式匹配
package leetcode

func isMatch(s string, p string) bool {
	if len(p) == 0 {
			if len(s) == 0 {
					return true
			} else {
					return false
			}
	}

	if len(s)==0 && len(p)==1{
			return false
	}

	m, n  := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
			dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 2; j <= n; j++ {
			if p[j-1] == '*' {
					dp[0][j] = dp[0][j-2]
			}
	}

	for i := 1; i <= m; i++ {
			for j:=1; j <= n; j++ {
					if s[i-1] == p[j-1] || p[j-1] == '.' {
							dp[i][j] = dp[i-1][j-1]
					} else if p[j-1] == '*' {
							if p[j-2] == s[i-1] || p[j-2] == '.' {
									dp[i][j] = dp[i-1][j] || dp[i][j-2]
							} else {
									dp[i][j] = dp[i][j-2]
							}
					} else {
							dp[i][j] = false
					}
			}
	}
	return dp[m][n]
}