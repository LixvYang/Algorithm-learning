// 验证二叉搜索树
package leetcode

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
			return true
	}
	if min != nil && root.Val <= min.Val {
			return false
	}
	if max != nil && root.Val >= max.Val {
			return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max) 
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
			if node == nil {
					return true
			}
			left := traverse(node.Left)
			if pre!=nil && node.Val <= pre.Val {
					return false
			}
			pre = node
			right := traverse(node.Right)
			return right&&left
	}

	return traverse(root)
}
