// 从前序与中序遍历序列构造二叉树
package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 && len(inorder) < 1 {
		return nil
	}

	nodeValue := preorder[0]
	left := findIndex(inorder, nodeValue)
	root := &TreeNode {
		Val: nodeValue,
		Left: buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}
	return root
}

func findIndex(inorder []int, nodeValue int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeValue {
			return i
		}
	}
	return -1
}