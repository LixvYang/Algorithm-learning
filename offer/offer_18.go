// 删除链表节点
func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}