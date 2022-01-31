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
 func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	RightMLeft(root, &sum)
	return root
}

func RightMLeft(root *TreeNode, sum *int)*TreeNode {
	if root == nil {
			return nil
	}
	RightMLeft(root.Right, sum)
	temp := *sum
	*sum+=root.Val
	root.Val+=temp
	RightMLeft(root.Left, sum)
	return root
}