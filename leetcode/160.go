// 相交链表
package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	pA, pB := headA, headB

	for pA != nil {
		pA = pA.Next
		lenA++
	}

	for pB != nil {
		pB = pB.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}