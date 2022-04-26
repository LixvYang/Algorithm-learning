package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

// 分治思想
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
			return nil
	}
	if len(lists) == 1{
			return lists[0]
	}
	mid := len(lists)/2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
					cur.Next = list1
					list1 = list1.Next
			} else {
					cur.Next = list2
					list2 = list2.Next
			}
			cur = cur.Next
	}

	if list1 != nil {
			cur.Next = list1
	} else if list2 != nil {
			cur.Next = list2
	}
	return head.Next
}