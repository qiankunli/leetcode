package linkedlist

import (
	"testing"
)

func TestRemoveNthFromEnd2(t *testing.T) {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	l := RemoveNthFromEnd2(l1, 2)
	Print(l)
}
