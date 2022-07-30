package linkedlist

func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	h := &ListNode{}
	h_t := h // 标记插入位置
	h_i := h // 标记链表尾部
	TraverseKList2(head, k, func(head, tail *ListNode, len int) {

		if len == k {
			InsertAfterByReverseOrder(h_i, head)
			h_t = head
			h_i = h_t
		} else {
			InsertAfterByOrder(h_i, head)
		}

	})
	return h.Next
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	h := &ListNode{}
	h_t := h // 标记插入位置
	h_i := h // 标记链表尾部
	TraverseList(head, func(i int, node *ListNode) {
		if i%k == 0 {
			h_i = h_t
			h_t = appendHead(h_i, node)
		} else {
			appendHead(h_i, node)
		}
	})
	return h.Next
}

// 把node 加在 insert 后面
func appendHead(insert, node *ListNode) *ListNode {
	insert_next := insert.Next
	h_c := &ListNode{Val: node.Val, Next: insert_next}
	insert.Next = h_c
	return h_c
}
