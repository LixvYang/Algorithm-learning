// Package leetcode provides 
package leetcode

import (
	"math"
	"container/list"
)

func largestValues(root *TreeNode) []int {
		res:=[][]int{}
    var finRes []int
    if root==nil{//防止为空
        return finRes
    }
    queue:=list.New()
    queue.PushBack(root)
    var tmpArr []int
    //层次遍历
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
		for i := 0; i < len(res); i++ {
			finRes = append(finRes, max(res[i]...))
		}
		return finRes
}

func max(vals...int) int {
	max := int(math.Inf(-1))
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}