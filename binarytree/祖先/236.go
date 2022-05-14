// 二叉树的公共祖先
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == p || root == q {
			return root
	}

	if root == nil {
			return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
			return nil
	}

	if left == nil && right != nil {
			return right
	}

	if left != nil && right == nil {
			return left
	}

	if left != nil && right != nil {
			return root
	}
	return nil
}