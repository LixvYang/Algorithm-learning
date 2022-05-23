package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
	if root == nil {
			return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if root.Left != nil && root.Right == nil {
			return 1+leftDepth
	}
	if root.Left == nil && root.Right != nil {
			return 1+rightDepth
	}
	return 1+min(leftDepth, rightDepth)
}

func min(a, b int) int {
	if a < b {
			return a
	}
	return b
}