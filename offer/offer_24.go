// 翻转链表
package offer

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
	}
	return pre
}