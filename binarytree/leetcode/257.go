// Package leetcode provides 
package leetcode


func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
			if node.Left == nil && node.Right == nil {
					v := s + strconv.Itoa(node.Val)
					res = append(res, v)
					return
			}
			s = s + strconv.Itoa(node.Val) + "->"
			if node.Left != nil {
					travel(node.Left, s)
			}
			if node.Right != nil {
					travel(node.Right, s)
			}
	}
	travel(root, "")
	return res
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			tmp := s+strconv.Itoa(node.Val)
			res = append(res, tmp)
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		travel(node.Left, s)
		travel(node.Right, s)
	}
	travel(root, "")
	return res
}