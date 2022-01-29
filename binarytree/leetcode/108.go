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
 func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0 {
			return nil
	}
	root := &TreeNode{nums[len(nums)/2],nil,nil}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right=sortedArrayToBST(nums[len(nums)/2+1:])
	return root
}