package leetcode

func fib(n int) int {
	// 边界条件
	if n < 2 {
		return n
	}
	
	// dp五部曲
	// 1. 定义dp数组，确定下标含义
	// 2. 递推公式 sum = dp[0] + dp[1] dp[0] = dp[1] dp[1] = sum
	// 3. dp初始化
	// 4. 确定遍历顺序
	// 5. 推举公式
	var sum int	
	dp := [2]int{}
	dp = [2]int{0, 1}

	for i:=1;i<n;i++ {
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}