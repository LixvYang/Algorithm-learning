// Package leetcode provides 
package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
			return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
			return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
			return lowestCommonAncestor(root.Right, p, q)
	} else {
			return root
	}
}