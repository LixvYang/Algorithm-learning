// Package leetcode provides 
package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}	
	res := 0
	if root.Right != nil {
		res+=countNodes(root.Right)
	}
	if root.Left != nil {
		res+=countNodes(root.Left)
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // labuladong题解
 // 先判断是否是子树是否有满二叉树  再进行计算
 func countNodes(root *TreeNode) int {
	l, r := root, root
	hl ,hr := 0, 0 
	for l != nil {
			l = l.Left
			hl++
	}
	for r != nil {
			r = r.Right
			hr++
	}
	if hl == hr {
			return int(math.Pow(2, float64(hl)))-1
	}
	return 1+countNodes(root.Left)+countNodes(root.Right)
}