// Package leetcode provides 
package leetcode

func minDepth(root *TreeNode) int {
	ans:=0
    if root==nil{
        return 0
    }
    queue:=list.New()
    queue.PushBack(root)
    for queue.Len()>0{
        length:=queue.Len()
        for i:=0;i<length;i++{
            node:=queue.Remove(queue.Front()).(*TreeNode)
            if node.Left==nil&&node.Right==nil{//当前节点没有左右节点，则代表此层是最小层
                return ans+1//返回当前层 ans代表是上一层
            }
            if node.Left!=nil{
                queue.PushBack(node.Left)
            }
            if node.Right!=nil{
                queue.PushBack(node.Right)
            }
        }
        ans++//记录层数
        

    }
    return ans+1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right != nil {
        return 1 + minDepth1(root.Right)
    }
    if root.Right == nil && root.Left != nil {
        return 1 + minDepth(root.Left);
    }
    return min(minDepth1(root.Left), minDepth1(root.Right))+1
}