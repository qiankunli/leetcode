package linkedlist

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	node3 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 3, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	l2 := &ListNode{Val: 1, Next: node4}

	node6 := &ListNode{Val: 6}
	l3 := &ListNode{Val: 2, Next: node6}

	lists := []*ListNode{l1, l2, l3}

	l := MergeKLists(lists)
	Print(l)

}
