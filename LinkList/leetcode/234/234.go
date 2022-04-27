package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
	}

	if fast != nil {
			slow = slow.Next
	}
	
	left := head
	right := reverseListNodes(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseListNodes(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
			next := head.Next
			head.Next = pre
			pre = head
			head = next
	}
	return pre
}