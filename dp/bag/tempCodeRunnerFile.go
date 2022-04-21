
// 	dp := make([][]int, len(weight))
// 	for i := range dp {
// 		dp[i] = make([]int, bagweight+1)
// 	}	

// 	for i := 0; i < len(weight); i++ {
// 		dp[i][0] = 0
// 	}

// 	for i := weight[0]; i <= bagweight; i++ {
// 		dp[0][i] = weight[0]
// 	}

// 	for i := 1; i < len(weight); i++ {
// 		for j := 1; j <= bagweight; j++ {
// 			if j < weight[i] {
// 				dp[i][j] = dp[i-1][j]
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+weight[i])
// 			}
// 		}
// 	}
// 	return dp[len(weight)-1][bagweight]
// }