package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
					cur.Next = l1
					cur = cur.Next
					l1 = l1.Next
			} else {
					cur.Next = l2
					cur = cur.Next
					l2 = l2.Next
			}
	}

	if l1 != nil {
			cur.Next = l1
	} else {
			cur.Next = l2
	}

	return head.Next
}   