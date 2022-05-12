// 把数字翻译成字符串
func translateNum(num int) int {
	if num < 10 {
	 return 1
	}
	str := strconv.Itoa(num)
	N := len(str)
	dp := make([]int, N+1)
	dp[0] = 1 // 这步需要反推
	dp[1] = 1

	for i := 2; i <= N; i++ {
		// 把前面两个数字合起来判断能否直接翻译(10-25之间都是能直接翻译的)
		temp := string(str[i-2]) + string(str[i-1]) 
		if temp >= "10" && temp <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	} 
	return dp[N]
}