// Package leetcode provides 
package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	var FindLeft func(node *TreeNode)
	FindLeft = func(node *TreeNode) {
			if node.Left != nil && node.Left.Left==nil && node.Left.Right==nil {
					res += node.Left.Val
			}
			if node.Left != nil {
					FindLeft(node.Left)
			}
			if node.Right != nil {
					FindLeft(node.Right)
			}
	}
	FindLeft(root)
	return res
}