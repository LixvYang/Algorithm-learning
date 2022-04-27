package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

 func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for curA != nil {
			curA = curA.Next
			lenA++
	}
	for curB!= nil {
			curB = curB.Next
			lenB++
	}
	var fast, slow *ListNode
	step := 0
	if lenA>lenB {
			fast = headA
			slow = headB
			step = lenA-lenB
	} else {
			fast = headB
			slow = headA
			step = lenB-lenA
	}
	for i:=0;i<step;i++ {
			fast = fast.Next
	}

	for fast!=slow {
			fast=fast.Next
			slow=slow.Next
	}
	return fast
}