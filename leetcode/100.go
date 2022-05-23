// 相同的二叉树
package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
			return true
	}

	if p == nil || q == nil {
			return false
	} 

	if p.Val != q.Val {
			return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
}