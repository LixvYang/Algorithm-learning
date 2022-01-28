// Package leetcode provides 
package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder)<1||len(postorder)<1{return nil}
	//先找到根节点（后续遍历的最后一个就是根节点）
	nodeValue:=postorder[len(postorder)-1]
	//从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	left:=findRootIndex(inorder,nodeValue)
	//构造root
	root:=&TreeNode{Val: nodeValue,
							Left: buildTree(inorder[:left],postorder[:left]),//将后续遍历一分为二，左边为左子树，右边为右子树
							Right: buildTree(inorder[left+1:],postorder[left:len(postorder)-1])}
	return root
}
func findRootIndex(inorder []int,target int) (index int){
	for i:=0;i<len(inorder);i++{
			if target==inorder[i]{
					return i
			}
	}
	return -1
}