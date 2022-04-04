// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var cur, head *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
				a = l1.Val
				l1 = l1.Next
		}

		if l2 != nil {
				b = l2.Val
				l2 = l2.Next
		}

		sum := a+b+carry
		sum, carry = sum%10, sum/10

		if head == nil {
				head = &ListNode{Val: sum}
				cur = head
		} else {
				cur.Next = &ListNode{Val:sum}
				cur = cur.Next
		}
	}

	if carry != 0 {
			cur.Next = &ListNode{Val:carry}
			cur = cur.Next
	}
	return head

}