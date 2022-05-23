// Package leetcode provides 
package leetcode

import (
	"container/list"
)

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr)
		tmpArr=[]int{}
	}

	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue = make([]*TreeNode, 0)
	queue = append(queue, root)
	if len(queue) > 0 {
		length := len(queue)
		tmpArr := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			tmpArr = append(tmpArr, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(ans, tmpArr)
	}
	return res
}