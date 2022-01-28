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
 func hasPathSum(root *TreeNode, targetSum int) bool {
	var flage bool
	if root == nil {
			return flage
	}

	pathsum(root,0,targetSum,&flage)
	return flage
}

func pathsum(root *TreeNode, sum int, targetSum int, flage *bool) {
	sum += root.Val
	if root.Left == nil &&root.Right == nil && sum == targetSum {
			*flage = true
			return 
	}
	if root.Left != nil {
			pathsum(root.Left, sum, targetSum, flage)
	}
	if root.Right != nil {
			pathsum(root.Right, sum, targetSum, flage)
	}
}