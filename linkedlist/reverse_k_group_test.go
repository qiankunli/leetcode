package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	l := ReverseKGroup(l1, 3)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
