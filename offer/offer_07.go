// 重建二叉树
package offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	nodeValue := preorder[0]
	left := findIndex(nodeValue, inorder)
	root := &TreeNode{
		Val: nodeValue,
		Left: buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}
	return root
}


func findIndex(target int, inorder []int) int {
	for i, v := range inorder {
			if v == target {
					return i
			}
	}
	return -1
}