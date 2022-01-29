package leetcode

// 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil||root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	searchBST(root.Right, val)
}