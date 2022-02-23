package bag

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
			sum+=v
	}
	target := sum/2
	dp := make([]int, 15001)

	for i:=0;i<len(stones);i++ {
			for j:=target;j>=stones[i];j-- {
					dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
	}
	return sum -2*dp[target]
}

func max(a,b int) int {
	if a<b {
			return b
	}
	return a
}