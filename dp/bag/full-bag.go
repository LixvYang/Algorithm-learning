package leetcode

func test_CompletePack() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4
	dp := make([]int, bagweight+1)
	// 先遍历物品 后遍历背包
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagweight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	// 先遍历背包后遍历物品
	for j := 0; j <= bagweight; j++ {
		for i := 0; i < len(weight); i++ {
			if j-weight[i] >= 0 {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
}