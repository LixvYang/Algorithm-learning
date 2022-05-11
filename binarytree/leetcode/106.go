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
 func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 && len(postorder) < 1 {
			return nil
	}

	node := postorder[len(postorder)-1]
	left := findIndex(node, inorder)

	root := &TreeNode{
			Val: node,
			Left: buildTree(inorder[:left], postorder[:left]),
			Right: buildTree(inorder[left+1:], postorder[left:len(postorder)-1]),
	}
	return root
}

func findIndex(val int, inorder []int) int {
	for i, v := range inorder {
			if v == val {
					return i
			}
	}
	return -1
}