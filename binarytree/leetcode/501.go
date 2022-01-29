// Package leetcod provides 
package leetcode

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode) 
	travel = func(node *TreeNode) {
			if node == nil {
					return
			}
			travel(node.Left)
			if prev != nil && prev.Val == node.Val {
					count++
			} else {
					count = 1
			}
			if count >= max {
					if count > max && len(res) > 0 {
							res = []int{node.Val}
					} else {
							res = append(res, node.Val)
					}
					max = count
			}
			prev = node
			travel(node.Right)
	}
	travel(root)
	return res
}