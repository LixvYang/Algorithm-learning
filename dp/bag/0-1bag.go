package main

import "fmt"

func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}

	for j := weight[0]; j <= bagweight; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		// for j := bagweight; j >= weight[i]; j-- {
		// 	dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		// }
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	for i := 0; i < len(weight); i++ {
		fmt.Printf("%d\n", dp[i])
	}
	return dp[len(weight)-1][bagweight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	println(test_2_wei_bag_problem1(weight, value, 4))
}

// func bag() {
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
