/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	var traverse func(root *TreeNode) bool
	traverse = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := traverse(root.Left)

		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		right := traverse(root.Right)
		return left && right
	}
	return traverse(root)
}
// @lc code=end

