package bag

func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}

	for j := weight[0]; j < bagweight+1; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return  dp[len(weight)-1][bagweight]

}

func max(a,b int) int {
	if a>b {
		return a 
	}
	return b
}

// func main() {
// 	weight := []int{1,3,4}
// 	value := []int{15,20,30}
// 	println(test_2_wei_bag_problem1(weight,value,4))
// }