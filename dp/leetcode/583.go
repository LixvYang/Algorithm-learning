package leetcode

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
for i := 0; i < len(dp); i++ {
	dp[i] = make([]int, len(word2)+1)
}
//初始化
for i := 0; i < len(dp); i++ {
	dp[i][0] = i
}
for j := 0; j < len(dp[0]); j++ {
	dp[0][j] = j
}
for i := 1; i < len(dp); i++ {
	for j := 1; j < len(dp[i]); j++ {
		if word1[i-1] == word2[j-1] {
			dp[i][j] = dp[i-1][j-1]
		} else {
			dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
		}
	}
}
return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
			return a 
	}
	return b
}