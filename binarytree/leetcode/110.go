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

func isBalanced(root *TreeNode) bool {
	return getHeigh(root) != -1
}

func getHeigh(root *TreeNode) int {
	if root == nil {
			return 0
	}

	leftH := getHeigh(root.Left)
	if leftH == -1 {
		return -1
	}
	rightH := getHeigh(root.Right)
	if rightH == -1 {
		return -1
	}

	if abs(leftH-rightH) > 1 {
		return -1
	} else {
		return max(rightH, leftH) + 1
	}
}
