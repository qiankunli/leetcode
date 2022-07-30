package linkedlist

import (
	"fmt"
	"testing"
)

func TestTraverseKList(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	TraverseKList(l1, 3, func(l *ListNode, len int) {
		fmt.Println("len ==> ", len)
		Print(l)
	})

}

func TestTraverseKList2(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	TraverseKList2(l1, 3, func(head, tail *ListNode, len int) {
		fmt.Println("len ==> ", len)
		fmt.Println("tail ==> ", tail.Val)
		Print(head)
	})

}

func TestInsertAfterByOrder(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	node6 := &ListNode{Val: 7}
	l2 := &ListNode{Val: 6, Next: node6}

	InsertAfterByOrder(l2, l1)
	Print(l2)
}

func TestInsertAfterByReverseOrder(t *testing.T) {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	node6 := &ListNode{Val: 7}
	l2 := &ListNode{Val: 6, Next: node6}

	InsertAfterByReverseOrder(l2, l1)
	Print(l2)
}
