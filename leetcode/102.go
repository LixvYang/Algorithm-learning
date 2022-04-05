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