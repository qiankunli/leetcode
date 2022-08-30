package linkedlist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	//node6 := &ListNode{Val: 1}
	//node5 := &ListNode{Val: 2, Next: node6}
	//node4 := &ListNode{Val: 3, Next: node5}
	//node3 := &ListNode{Val: 3, Next: node4}
	//node2 := &ListNode{Val: 2, Next: node3}
	//l1 := &ListNode{Val: 1, Next: node2}
	//
	//fmt.Println(IsPalindrome(l1))

	node4 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2, Next: node4}
	node2 := &ListNode{Val: 1, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	fmt.Println(IsPalindrome(l1))
}
