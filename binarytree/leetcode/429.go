// Package leetcode provides 
package leetcode

import (
	"container/list"
)

type Node struct {
	Val int
	Children []*Node
}


func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for T := 0; T < length; T++ {
			myNode := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res, tmp)
	}
	return res
}