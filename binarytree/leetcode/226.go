// Package leetcode provides 
//反转二叉树

package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)
	
	return root
}