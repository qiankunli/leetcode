package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	node5 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 3, Next: node5}
	l2 := &ListNode{Val: 1, Next: node4}

	l := MergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
