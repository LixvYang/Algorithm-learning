// 打家劫舍3
package leetcode

func rob(root *TreeNode) int {
	dp := traversal(root)

	return max(dp[0], dp[1])
}

func traversal(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	dpL := traversal(root.Left)
	dpR := traversal(root.Right)

	val1 := root.Val + dpL[0] + dpR[0]
	val2 := max(dpL[0], dpL[1]) + max(dpR[0], dpR[1])
	return []int{val1, val2}
}