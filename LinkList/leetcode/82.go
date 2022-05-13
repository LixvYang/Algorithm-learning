// 删除链表的重复元素
package leetcode 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
					return head
			}
			dummy := &ListNode{Next:head}
			prev, cur := dummy, head
			for cur != nil{
					for cur.Next != nil && cur.Val == cur.Next.Val{
							cur = cur.Next
					}
					if prev.Next == cur{
							prev = prev.Next
					}else{
							prev.Next = cur.Next
					}
					cur = cur.Next
			}
			return dummy.Next  
	
	}
	
	