// Package leetcode provides 
package leetcode

import (
	"container/list"
)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func averageOfLevels(root *TreeNode) []float64 {
	res:=[][]int{}
	var finRes []float64
	if root==nil{//防止为空
			return finRes
	}
	queue:=list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len()>0 {
			length:=queue.Len()//保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
			for i:=0;i<length;i++{
					node:=queue.Remove(queue.Front()).(*TreeNode)//出队列
					if node.Left!=nil{
							queue.PushBack(node.Left)
					}
					if node.Right!=nil{
							queue.PushBack(node.Right)
					}
					tmpArr=append(tmpArr,node.Val)//将值加入本层切片中
			}
			res=append(res,tmpArr)//放入结果集
			tmpArr=[]int{}//清空层的数据
	}
	
	length := len(res)
	for i := 0; i < length; i++ {
			var sum int
			for j := 0; j < len(res[i]); j++ {
					sum += res[i][j]
			}
			tmp := float64(sum)/float64(len(res[i]))
			finRes = append(finRes, tmp)
	}
	return finRes
}