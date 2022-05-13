// 删除BST中的节点

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

 func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
			return nil
	}
	if root.Val == key {
			if root.Left == nil && root.Right == nil {
					root = nil
					return nil
			}
			if root.Left != nil && root.Right == nil {
					root = root.Left
					return root
			}
			if root.Right != nil && root.Left == nil {
					root = root.Right
					return root
			}
			if root.Right != nil && root.Left != nil {
					left := root.Left
					right := root.Right
					tmp := root.Right
					for tmp.Left != nil {
							tmp = tmp.Left
					}
					tmp.Left = left
					root = right
					return root
			}
	}

	if root.Val < key {
			root.Right = deleteNode(root.Right, key)
	}
	if root.Val > key {
			root.Left = deleteNode(root.Left, key)
	}
	return root
}