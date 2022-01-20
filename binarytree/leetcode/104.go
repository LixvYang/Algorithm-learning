// Package leetcode provides 
package leetcode

import (
	"container/list"
)

func maxDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
			return ans
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
			length := queue.Len()
			for i:=0;i<length;i++{
					node:=queue.Remove(queue.Front()).(*TreeNode)
					if node.Left!=nil{
							queue.PushBack(node.Left)
					}
					if node.Right!=nil{
							queue.PushBack(node.Right)
					}
			}
			ans++
	}
	return ans
}