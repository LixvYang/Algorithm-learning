// 二叉树展开为链表
package leetcode

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil 
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}