package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	l := Reverse(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
