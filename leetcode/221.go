// 最大正方形
package leetcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = n+1
	}
	max := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))+1
			}
			
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max*max
}