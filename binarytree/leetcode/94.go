package leetcode

func  inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode)  {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}