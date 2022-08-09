package linkedlist

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}
	fmt.Println(HasCycle(l1))

	node := &ListNode{Val: 4}
	node8 := &ListNode{Val: 4, Next: node}
	node7 := &ListNode{Val: 3, Next: node8}
	node.Next = node7
	node6 := &ListNode{Val: 2, Next: node}
	l2 := &ListNode{Val: 1, Next: node6}
	fmt.Println(HasCycle(l2))
}
