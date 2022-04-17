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

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
	res:=[][]*Node{}
    if root == nil {
        return root
    }
    queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmpArr := make([]*Node, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, node)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmpArr)
	}

    for i := 0; i < len(res); i++ {
        for j := 0; j < len(res[i])-1; j++ {
            res[i][j].Next = res[i][j+1]
        }
    }
    return root
}