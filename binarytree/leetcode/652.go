// 寻找重复自树
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var mapping  map[string]int //记录重复子树
 var nodeList []*TreeNode //记录重复结点
 
	func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
		 mapping = make(map[string]int)
		 nodeList = make([]*TreeNode,0)
		 traverse(root)
		 return nodeList
 }
 
 func traverse(root *TreeNode)string{
		 if root == nil{
				 return "#"
		 }
		 left :=  traverse(root.Left)
		 right := traverse(root.Right)
 
		 subTree := left + "#" + right +"#"+ strconv.Itoa(root.Val)
		 if num,ok := mapping[subTree];ok{
				 if num == 1{
						 nodeList = append(nodeList,root)
						 mapping[subTree]++
				 }
		 }else{
					mapping[subTree]++
		 }
		 
		 return subTree
 }
