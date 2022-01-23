// Package leetcode provides 
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
	if root == nil {
			return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
			return false
	}

	LeftH := maxdepth(root.Left)
	RightH := maxdepth(root.Right)

	if abs(LeftH-RightH) > 1 {
			return false
	}
	return true
}

func maxdepth(root *TreeNode) int {
	if root == nil {
			return 0
	}

	return max(maxdepth(root.Left),maxdepth(root.Right))+1
}

func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
			return -a
	}
	return a
}