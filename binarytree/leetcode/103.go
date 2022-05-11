package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	flag := true
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
							queue = append(queue, node.Right)
					}
			}
			if !flag {
					reverse(tmpArr)
			}
			flag = !flag
			res = append(res, tmpArr)
	}
	return res
}   

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
	}
}