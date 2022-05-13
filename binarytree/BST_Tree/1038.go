// 1038. 从二叉搜索树到更大和树
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstToGst(root *TreeNode) *TreeNode {
	var pre int
	traverse(root, &pre)
	return root
}

func traverse(root *TreeNode, pre *int) {
	if root == nil {
			return
	}
	traverse(root.Right, pre)
	root.Val += *pre
	*pre = root.Val
	traverse(root.Left, pre)
}