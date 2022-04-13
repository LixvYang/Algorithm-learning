//  链表中倒数第k个节点
package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre, cur := dummy, head

	for k>0 {
			cur = cur.Next
			k--
	}
	
	if cur == nil {
			return head
	}

	for cur != nil {
			cur = cur.Next
			pre = pre.Next
	}

	return pre.Next
}
