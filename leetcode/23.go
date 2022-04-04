// 合并K个升序链表
package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
			return nil
	}
	if len(lists) == 1{
			return lists[0]
	}

	var result *ListNode
	for _, list := range lists {
			result = mergeTwoLists(result, list)
	}
	return result
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