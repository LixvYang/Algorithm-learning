package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i:=0;i<m && obstacleGrid[i][0] == 0;i++ {
		dp[i][0] = 1
	}
	for i:=0;i<n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i:=0;i<m;i++ {
		for j:=0; j<n;j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}	