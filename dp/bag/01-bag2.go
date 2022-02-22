package main

func test_1_wei_bag_problem(weight, value []int, bagweight int) int {
	dp := make([]int, bagweight+1)

	// 递推顺序
	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagweight]
}

func max(a,b int) int {
	if a>b {return a}
	return b
}

func main() {
	weight := []int{1,3,4}
	value := []int{15,20,30}
	println(test_1_wei_bag_problem(weight,value,4))
}