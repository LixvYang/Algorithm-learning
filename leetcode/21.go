// 合并两个有序链表
package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTowLists(list1.Next, list2)
		return l1
	} else {
		list2.Next = mergeTowLists(list2.Next, list1)
		return l2
	}
}

func mergeTowLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l1 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		}  else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}
