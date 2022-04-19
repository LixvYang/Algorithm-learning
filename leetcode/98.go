// 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/
package leetcode

func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return 
		}

		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)

	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		traverse(node.Left)
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		traverse(node.Right)
	}
	return traverse(root)
}