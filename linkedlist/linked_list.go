package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertAfter(cur, node *ListNode) {
	next := cur.Next
	cur.Next = node
	node.Next = next
}
func InsertAfterByOrder(cur, l *ListNode) {
	next := cur.Next
	TraverseList(l, func(i int, node *ListNode) {
		cur.Next = node
		cur = cur.Next
	})
	cur.Next = next
}
func InsertAfterByReverseOrder(cur, l *ListNode) {
	next := cur.Next
	TraverseList(l, func(i int, node *ListNode) {
		InsertAfter(cur, node)
	})
	l.Next = next
}
func TraverseList(head *ListNode, handler func(i int, node *ListNode)) {
	cur := head
	i := 0
	for cur != nil {
		next := cur.Next // node 可以被handler 挪用
		handler(i, cur)
		i++
		cur = next
	}
}
func TraverseListByNode(head *ListNode, handler func(i int, isTail bool, node *ListNode)) {
	cur := head
	i := 0
	for cur != nil {
		next := cur.Next // node 可以被handler 挪用
		cur.Next = nil   // node 的后缀关系也被清除
		handler(i, nil == next, cur)
		i++
		cur = next
	}
}
func TraverseListByValue(head *ListNode, handler func(i, value int)) {
	cur := head
	i := 0
	for cur != nil {
		handler(i, cur.Val)
		i++
		cur = cur.Next
	}
}

func TraverseKList(head *ListNode, k int, handler func(l *ListNode, len int)) {

	l := &ListNode{}
	cur := l
	count := 0
	TraverseListByNode(head, func(i int, isTail bool, node *ListNode) {
		cur.Next = node
		cur = cur.Next
		count++
		if (i+1)%k == 0 || isTail {
			handler(l.Next, count)
			l = &ListNode{}
			cur = l
			count = 0
		}
	})
}

func TraverseKList2(head *ListNode, k int, handler func(head, tail *ListNode, len int)) {

	l := &ListNode{}
	cur := l
	count := 0
	TraverseListByNode(head, func(i int, isTail bool, node *ListNode) {
		cur.Next = node
		cur = cur.Next
		count++
		if (i+1)%k == 0 || isTail {
			handler(l.Next, cur, count)
			l = &ListNode{}
			cur = l
			count = 0
		}
	})
}

func Print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
