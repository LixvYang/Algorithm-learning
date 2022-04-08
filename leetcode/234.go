// 回文链表
package leetcode

func isPalindrome(head *ListNode) bool {
	vals := []int{}

	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	start, end := 0, len(vals)-1
	for start < end {
		if vals[start] != vals[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// 分为两个链表
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	var preSlow *ListNode
	for slow != nil {
		preSlow = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	pre.Next = nil // 断开
	// 反转
	var head2 *ListNode = nil
	for slow != nil {
		temp := slow.Next
		slow.Next = head2
		head2 = slow
		slow = temp
	}

	// 比对
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2= head2.Next
	}
	return true
}