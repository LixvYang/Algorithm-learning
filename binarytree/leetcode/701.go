// Package leetcode provides 
package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
			root = &TreeNode{Val:val}
			return root
	}

	if root.Val > val {
			root.Left = insertIntoBST(root.Left, val)
	} else {
			root.Right = insertIntoBST(root.Right, val)
	}
	return root
}