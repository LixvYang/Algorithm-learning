// Package leetcode provides 
package leetcode

import (
	"container/list"
)

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode)  {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode)  {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans

}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		ans = append(ans, node.Val)
		stack = stack[:len(stack)-1]
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return ansgi
}