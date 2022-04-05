// 对称二叉树
package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Compare(root.Left, root.Right)
}

func Compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		returnt true
	}

	if left = nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return Compare(left.Left, right.Right) && Compare(left.Right, right.Left)
}

