//230. 二叉搜索树中第K小的元素
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
			if root == nil {
					return
			}
			traverse(root.Left)
			res = append(res, root.Val)
			traverse(root.Right)
	} 
	traverse(root)
	return res[k-1]
}