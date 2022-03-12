package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
t2 := len(text2)
dp:=make([][]int,t1+1)
for i:=range dp{
	dp[i]=make([]int,t2+1)
}

	for i:=1;i<=t1;i++ {
			for j:=1;j<=t2;j++ {
					if text1[i-1] == text2[j-1] {
							dp[i][j] = dp[i-1][j-1]+1
					}else {
							dp[i][j] = max(dp[i-1][j],dp[i][j-1])
					}
			}
	}
	return dp[t1][t2]
}

func max(a, b int) int {
	if a < b {
			return b
	}
	return a
}