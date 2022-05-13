package offer

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
			return res
	}
	currPath := make([]int, 0)
	currPath = append(currPath, root.Val)
	traverse(root, currPath, &res, targetSum-root.Val)
	return res
}

func traverse(node *TreeNode, currPath []int, res *[][]int, count int) {
	if node.Left == nil && node.Right == nil {
			if count == 0 {
					temp := make([]int, len(currPath))
					copy(temp, currPath)
					*res = append(*res,temp)
			} else {
					return
			}
	}

	if node.Left != nil {
			currPath = append(currPath, node.Left.Val)
			count -= node.Left.Val
			traverse(node.Left, currPath, res, count)
			count += node.Left.Val
			currPath = currPath[:len(currPath)-1]
	}

	if node.Right != nil {
			currPath = append(currPath, node.Right.Val)
			count -= node.Right.Val
			traverse(node.Right, currPath, res, count)
			count += node.Right.Val
			currPath = currPath[:len(currPath)-1]
	}
	return
}