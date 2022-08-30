package linkedlist

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	//node3 := &ListNode{Val: 6}
	//node2 := &ListNode{Val: 1, Next: node3}
	//l1 := &ListNode{Val: 7, Next: node2}
	//
	//node5 := &ListNode{Val: 2}
	//node4 := &ListNode{Val: 9, Next: node5}
	//l2 := &ListNode{Val: 5, Next: node4}
	//
	//r := AddTwoNumbers(l1, l2)
	//Print(r)

	node2 := &ListNode{Val: 8}
	l1 := &ListNode{Val: 1, Next: node2}

	l2 := &ListNode{Val: 0}

	r := AddTwoNumbers(l1, l2)
	Print(r)
}
