package bag

func test_CompletePack1(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	dp[0] = 1

	for i:=0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

