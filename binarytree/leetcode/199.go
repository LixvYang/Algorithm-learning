package leetcode

import (
	"container/list"
)
/**
199. 二叉树的右视图
 */
 func rightSideView(root *TreeNode) []int {
	queue:=list.New()
	res:=[][]int{}
	var finaRes []int
	if root==nil{
			return finaRes
	}
	queue.PushBack(root)
	for queue.Len()>0{
			length:=queue.Len()
			tmp:=[]int{}
			for i:=0;i<length;i++{
					node:=queue.Remove(queue.Front()).(*TreeNode)
					if node.Left!=nil{
							queue.PushBack(node.Left)
					}
					if node.Right!=nil{
							queue.PushBack(node.Right)
					}
					tmp=append(tmp,node.Val)
			}
			res=append(res,tmp)
	}
	//取每一层的最后一个元素
	for i:=0;i<len(res);i++{
			finaRes=append(finaRes,res[i][len(res[i])-1])
	}
	return finaRes
}

func rightSideView(root *TreeNode) []int {
	res := make([][]int, 0)
	Right := make([]int, 0)
	if root == nil {
		return Right
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmpArr := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmpArr)
	}
	for i := 0; i < len(res); i++ {
		Right = append(Right, res[i][len(res[i])-1])
	}
	return Right
}