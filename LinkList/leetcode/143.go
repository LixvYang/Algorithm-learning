// 重拍链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return	
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
	}

	head1 := head
	head2 := reverseList(slow)

	for head1 != head2 && head1.Next != head2 {
	next1, next2 := head1.Next, head2.Next
	head1.Next, head2.Next = head2, next1
	head1, head2 = next1, next2
}

}

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
