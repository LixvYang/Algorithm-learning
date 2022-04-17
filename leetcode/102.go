// 二叉树的层序遍历
package leetcode

import "list"

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmpArr := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, tmpArr)
	}
	return res
}


func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmpArr := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue.append(queue, node.Right)
			}
		}
		res = append(res, tmpArr)
	}
	return res
}