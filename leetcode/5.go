// 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
// Language: go
// Path: leetcode\5.go
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	} 

	var res string
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s)-1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <=1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}

			if dp[i][j] && j-i+1 > len(res) {
				res = s[i:j+1]
			}
		}
	}
}