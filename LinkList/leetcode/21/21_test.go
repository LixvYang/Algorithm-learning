package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l3 := &ListNode{6, nil}
	l2 := &ListNode{4, l3}
	l1 := &ListNode{2, l2}

	n3 := &ListNode{5, nil}
	n2 := &ListNode{3, n3}
	n1 := &ListNode{1, n2}
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{	
		{l1, n1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}},g
	}

	for _, test := range tests {
		got := MergeTwoLists(test.list1, test.list2)
		if got != test.want {
			t.Errorf("mergeTwoLists(%v, %v) = %v; want %v", test.list1, test.list2, got, test.want)
		}
	}
}
