// Package leetcode provides 
package leetcode

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 递归删除节点, 找到节点，递补过去
// 5种情况：
// 1.空节点返回nil
// 2.左右孩子都为空，直接删除节点返回nil
// 3.左空右不空，删除节点，右孩子补位
// 4.左不空右空，删除节点，左孩子补位
// 5.左右孩子节点都不为空，则将删除节点的左子树放到删除节点的右子树的最左面节点的左孩子的位置
if root == nil { // 1.空节点返回nil
	return nil
}
if root.Val == key {
	// 2.左右孩子都为空，直接删除节点返回nil
	if root.Right == nil && root.Left == nil {
		return nil
	}
	// 3.左空右不空，删除节点，右孩子补位
	if root.Left == nil && root.Right != nil {
		root = root.Right
		return root
	}
	// 4.左不空右空，删除节点，左孩子补位
	if root.Left != nil && root.Right == nil {
		root = root.Left
		return root
	}
	// 5.左右孩子节点都不为空，则将删除节点的左子树放到删除节点的右子树的最左面节点的左孩子的位置
	// root为删除节点，保存左孩子树
	left := root.Left
	// 保存右孩子树
	right := root.Right
	// 通过tmp不断遍历到root.Right的最左孩子节点
	tmp := root.Right
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	// 右子树的最小节点，该节点大于root.Left整颗子树，Left指向Left子树
	tmp.Left = left
	// 更新root节点，相当于删除root节点
	root = right
	return root
}
if root.Val > key { // 递归左树
	root.Left = deleteNode(root.Left, key)
}
if root.Val < key { // 递归右树
	root.Right = deleteNode(root.Right, key)
}
return root
}