package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p1 != nil && p1.Next != nil {
			p1 = p1.Next.Next
			p2 = p2.Next
	}
	return p2
}