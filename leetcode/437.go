// 路径总和3
type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
			return
	}
	val := root.Val
	if val == targetSum {
			res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
			return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
