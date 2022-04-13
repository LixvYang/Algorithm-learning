// 从尾到头打印链表
package offer

func reversePrint(head *ListNode) []int {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	result := make([]int, 0)
	for pre != nil {
		result = append(result, pre.Val)
		pre = pre.Next
	}
	return result
}