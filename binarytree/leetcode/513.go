// Package leetcode provides 
package leetcode

func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
			return root.Val
	}

	var maxDeep, value int
	var findLeftValue func(root *TreeNode, deep int)
	findLeftValue = func(root *TreeNode, deep int) {
			if root.Left == nil && root.Right == nil {
					if deep > maxDeep {
							value = root.Val
							maxDeep = deep
					}
			}
			
			if root.Left != nil {
					deep++
					findLeftValue(root.Left, deep)
					deep--
			}

			if root.Right != nil {
					deep++
					findLeftValue(root.Right, deep)
					deep--
			}
	}
	findLeftValue(root, maxDeep)
	return value

}