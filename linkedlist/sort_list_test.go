package linkedlist

import "testing"

func TestSortList(t *testing.T) {

	node3 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}
	Print(SortList(l1))
}
