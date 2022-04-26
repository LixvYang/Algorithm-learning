package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}



func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next:head}
	x := findFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre, cur := dummy, head

	// i := 1
	for n>0 {
			cur = cur.Next
			n--
	}
	for cur != nil {
			cur = cur.Next
			pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return dummy.Next
}


// 寻找第倒数n个的节点的值
func findFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head

	for n > 0 {
			p1 = p1.Next
			n--
	}

	for p1 != nil {
			p1 = p1.Next
			p2 = p2.Next
	}
	return p2
}