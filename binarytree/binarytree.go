
package main

import (
   "fmt"
)
//定义二叉树结构体
type binaryTree struct {
 	value int
  leftNode *binaryTree
  rightNode *binaryTree
}

func main(){
   //0代表节点为空
  num := []int{1,2,3,4,5,0,6}
  //将数组切片转化为二叉树结构体
  tree := createBinaryTree(0,num)
   //二叉树数据
  var front []int
  front = frontForeach(*tree,front)
   fmt.Println(front)
   fmt.Print("leftNode",tree.rightNode)
}

//创建二叉树
func createBinaryTree(i int,nums []int) *binaryTree {
   tree := &binaryTree{nums[i], nil, nil}
   //左节点的数组下标为1,3,5...2*i+1
  if i<len(nums) && 2*i+1 < len(nums) {
      tree.leftNode = createBinaryTree(2*i+1, nums)
   }
   //右节点的数组下标为2,4,6...2*i+2
  if i<len(nums) && 2*i+2 < len(nums) {
      tree.rightNode = createBinaryTree(2*i+2, nums)
   }
   return tree
}

//前序遍历
func frontForeach(tree binaryTree,num []int) []int{
   //先遍历根节点
  if tree.value != 0 {
      num = append(num,tree.value)
   }
   var leftNum,rightNum []int
  //若存在左节点，遍历左节点树
  if tree.leftNode != nil {
      leftNum = frontForeach(*tree.leftNode,leftNum)
      for _,value := range leftNum{
         num = append(num,value)
      }
   }
   //若存在右节点，遍历右节点树
  if tree.rightNode != nil {
      rightNum = frontForeach(*tree.rightNode,rightNum)
      for _,value := range rightNum{
         num = append(num,value)
      }
   }

   return num
}