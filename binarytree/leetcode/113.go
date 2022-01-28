// Package leetcode provides 
package leetcode

func pathsum(root *TreeNode, targetSum int) [][]int {
	var result [][]int	
	if root == nil {
		return result
	}

	var sumnodes []int
	haspathsum(root, &sumnodes, targetSum, &result)
	return result
}

func haspathsum(root *TreeNode, sumnodes *[]int, targetSum int, result *[][]int)  {
	*sumnodes=append(*sumnodes, root.Val)

	if root.Left == nil && root.Right == nil {
		var sum, number int
		for k, v := range *sumnodes {
			sum += v
			number = k
		}

		tempnodes:= make([]int, number+1)
		for k, v := range *sumnodes {
			tempnodes[k]=v
		}
		if sum == targetSum {
			*result = append(*result, tempnodes)
		}
	}
	if root.Left != nil {
		haspathsum(root.Left, sumnodes, targetSum, result)
		*sumnodes = (*sumnodes)[:len(*sumnodes)-1]
	}
	if root.Right != nil {
		haspathsum(root.Right, sumnodes, targetSum, result)
		*sumnodes = (*sumnodes)[:len(*sumnodes)-1]
	}
}