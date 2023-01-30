/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode.cn/problems/integer-break/description/
 *
 * algorithms
 * Medium (62.34%)
 * Likes:    1030
 * Dislikes: 0
 * Total Accepted:    225.7K
 * Total Submissions: 362.1K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
 * 
 * 返回 你可以获得的最大乘积 。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: n = 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 * 
 * 示例 2:
 * 
 * 
 * 输入: n = 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 2 <= n <= 58
 * 
 * 
 */

// @lc code=start
func integerBreak(n int) int {
	if n < 3 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

