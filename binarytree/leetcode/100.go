// Package leetcode provides 
package leetcode

// 求相同的树

func isSamepleTree(p *TreeNode, q *TreeNode) bool {
	switch  {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil :
		fallthrough
	case p.Val != q.Val:
		return false
	}
	return isSamepleTree(p.Left, q.Left) && isSamepleTree(p.Right, q.Right)
}