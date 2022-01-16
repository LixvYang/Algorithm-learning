// Package leetcode provides 
package leetcode

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode)  {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}