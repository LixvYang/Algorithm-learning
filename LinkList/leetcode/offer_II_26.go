package leetcode

type ListNode struct {
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
 return
}

 fast, slow := head, head
 for fast != nil && fast.Next != nil {
		 fast = fast.Next.Next
		 slow = slow.Next
 }

 head1 := head
 head2 := reverse(slow)

 for head1 != head2 && head1.Next != head2 {
 next1, next2 := head1.Next, head2.Next
 head1.Next, head2.Next = head2, next1
 head1, head2 = next1, next2
}
}

func reverse(head *ListNode) *ListNode {
 var pre *ListNode
 for head != nil {
		 nxt := head.Next
		 head.Next = pre
		 pre = head
		 head = nxt
 }
 return pre
}