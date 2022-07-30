package linkedlist

func Reverse(head *ListNode) *ListNode {
	return reverse2(head)
}

// 反转链表
func reverse1(head *ListNode) *ListNode {
	h := &ListNode{}
	cur := head
	for cur != nil {
		h_next := h.Next
		h_c := &ListNode{Val: cur.Val, Next: h_next}
		h.Next = h_c
		cur = cur.Next
	}
	return h.Next
}

// 反转链表
func reverse2(head *ListNode) *ListNode {
	h := &ListNode{}
	TraverseList(head, func(i int, node *ListNode) {
		InsertAfter(h, node)
	})
	return h.Next
}
