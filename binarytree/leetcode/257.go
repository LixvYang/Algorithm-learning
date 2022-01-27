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