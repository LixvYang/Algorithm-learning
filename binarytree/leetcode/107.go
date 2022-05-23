package leetcode

import (
	"container/list"
)

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmpArr := []int{}
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
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
