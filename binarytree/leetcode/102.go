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