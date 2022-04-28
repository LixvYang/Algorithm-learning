package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
			return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last  
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	if head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}


var successor *ListNode
//反转链表前N个元素
func reverseN(head *ListNode, n int) *ListNode {	
	if n == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}


// 反转链表的一部分
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	
	between := reverseBetween(head.Next, m-1, n-1)
	head.Next = between
	return head
}