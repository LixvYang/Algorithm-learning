// Package leetcode provides 
package leetcode

func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMin(root, &res)
	min := 10000000
	for i := 1; i < len(res); i++ {
			tempValue := res[i]-res[i-1]
			if tempValue < min {
					min = tempValue
			}
			
	}
	return min
}

func findMin(root *TreeNode, res *[]int) {
	if root == nil {
			return
	}
	findMin(root.Left, res)
	*res = append(*res, root.Val)
	findMin(root.Right, res)
}

func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt32
	var pre *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return 
		}
		dfs(node.Left)
		if pre != nil {
			result = min(result, node.Val - pre.Val)
		}
		pre = node

		dfs(node.Right)
	}

	dfs(root)
	return result
}
