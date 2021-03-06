// Package leetcod provides 
package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)<1{return nil}
	//首选找到最大值
	index:=findMax(nums)
	//其次构造二叉树
	root:=&TreeNode{
			Val: nums[index],
			Left:constructMaximumBinaryTree(nums[:index]),//左半边
			Right:constructMaximumBinaryTree(nums[index+1:]),//右半边
			}
	return root
}
func findMax(nums []int) (index int){
	for i:=0;i<len(nums);i++{
			if nums[i]>nums[index]{
					index=i
			}
	}
	return 
}