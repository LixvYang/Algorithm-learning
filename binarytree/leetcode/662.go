// 二叉树最大宽度
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	mp := map[*TreeNode]int{}
	mp[root] = 1
	res := 1
	for len(queue) > 0 {
			length := len(queue)
			width := mp[queue[len(queue)-1]] - mp[queue[0]]+1
			if width > res {
					res = width
			}
			for i := 0; i < length; i++ {
					node := queue[0]
					queue = queue[1:]
					if node.Left != nil {
							queue = append(queue, node.Left)
							mp[node.Left] = mp[node]*2
					}
					if node.Right != nil {
							queue = append(queue, node.Right)
							mp[node.Right] = mp[node]*2+1
					 }
			}
	}
	return res
}

	
