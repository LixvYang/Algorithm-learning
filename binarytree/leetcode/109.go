package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    arr := make([]int, 0)
    for head != nil {
        arr =append(arr, head.Val)
        head = head.Next
    }
    return buildTree(arr, 0, len(arr)-1)
}

func buildTree(arr []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }
    mid := (start+end) >> 1
		root := &TreeNode{Val: arr[mid]}
		root.Left = buildBST(arr, start, mid-1)
		root.Right = buildBST(arr, mid+1, end)
		return root
    return root
}