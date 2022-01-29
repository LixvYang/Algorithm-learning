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