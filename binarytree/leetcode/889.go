// packag leetcode
package leetcode

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) < 1 && len(postorder) < 1 {
		return nil
	}

	node := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{Val: node}
	}

	idx := 0
	for i, v := range postorder {
		if v == preorder[1] {
			idx = i+1
			break
		}
	}

	root := &TreeNode{
		Val: node,
		Left: constructFromPrePost(preorder[1:idx+1], postorder[:idx]),
		Right: constructFromPrePost(preorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
	return root
}