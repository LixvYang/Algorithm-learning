package test

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			tmp = append(tmp, node.Value)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}