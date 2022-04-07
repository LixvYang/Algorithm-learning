// 排序链表
package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	preSlow := &ListNode{}
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil
	l := sortList(head)
	r := sortList(slow)
	return mergeList(l, r)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
			pre = pre.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
			pre = pre.Next
		}
	}

	if l1 != nil {
		pre.Next = l1
	}

	if l2 != nil {
		pre.Next = l2
	}

	return dummy.Next
}