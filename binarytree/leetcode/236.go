// Package leetcode provides 
package leetcode

// 二叉树的公共祖先问题
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
			return root
	}
	
	if root == p || root == q {
			return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
			return root
	} else if left != nil {
			return left
	} else {
			return right
	}
	return nil
}