/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode.cn/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (65.88%)
 * Likes:    1589
 * Dislikes: 0
 * Total Accepted:    365.4K
 * Total Submissions: 554.5K
 * Testcase Example:  '12'
 *
 * 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
 * 
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 12
 * 输出：3 
 * 解释：12 = 4 + 4 + 4
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^4
 * 
 * 
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}

	// 遍历物品
	for i := 1; i <= n; i++ {
		// 遍历背包
		for j := i*i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if  a < b {
		return a 
	}
	return b
}
// @lc code=end

