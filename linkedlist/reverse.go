package linkedlist

// 206
func Reverse(head *ListNode) *ListNode {
	return reverse3(head)
}

// 反转链表

/*
	将链表断为两段，一段转移到另一段
	1 -> 2 -> 3 -> 4
	1 -> nil   2-> 3-> 4
	2 -> 1 -> nil   3 -> 4
*/
func reverse4(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode // 新链表的头节点
	cur := head       // 原有链表的当前节点
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*
	head -> 1 -> 2 -> 3 -> 4
	head -> 2 -> 1 -> 3 -> 4
	head -> 3 -> 2 -> 1 -> 4
	head -> 4 -> 3 -> 2 -> 1
*/
// 伪头指针
// 头插法
func reverse3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 只有一个元素
	if head.Next == nil {
		return head
	}
	h := &ListNode{Next: head} // 虚拟头节点
	cur := head                // cur 的下一个元素为待移动的节点
	for cur.Next != nil {

		next := cur.Next
		// 将next 抽出  顺带右移cur
		cur.Next = next.Next

		// 将next 移到h 后
		next.Next = h.Next
		h.Next = next
	}
	return h.Next
}

func reverse1(head *ListNode) *ListNode {
	h := &ListNode{}
	cur := head
	for cur != nil {
		h_next := h.Next
		h_c := &ListNode{Val: cur.Val, Next: h_next}
		h.Next = h_c
		cur = cur.Next
	}
	return head.Next
}

// 反转链表
func reverse2(head *ListNode) *ListNode {
	h := &ListNode{}
	TraverseList(head, func(i int, node *ListNode) {
		InsertAfter(h, node)
	})
	return h.Next
}
