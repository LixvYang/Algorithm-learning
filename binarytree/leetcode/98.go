// Package leetcode provides 
package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return check(root, math.MinInt64,math.MaxInt64)
}

func check(node *TreeNode, min,max int64) bool {
    if node == nil {
        return true
    }
    left := check(node.Left,min,int64(node.Val))
    if min >= int64(node.Val) || max <= int64(node.Val) {
        return false
    }
    right := check(node.Right,int64(node.Val),max)
    return left && right
}