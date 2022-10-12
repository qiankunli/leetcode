package linkedlist

// 61
func rotateRight(head *ListNode, k int) *ListNode {
	// 如果链表长度大于k，则转换为寻找第k个节点，将后面k个节点 接到链表的头部
	// 如果链表长度小于k，则转换为寻找第k-n个节点

	n := Len(head)
	if n == 0 {
		return head
	}
	if k > n {
		k = k % n
	}
	if k == n {
		return head
	}

	dummy := &ListNode{Next: head}
	p := dummy  // 指向最后一个节点
	kp := dummy // 第k个节点前一个节点

	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p.Next != nil {
		p = p.Next
		kp = kp.Next
	}

	p.Next = dummy.Next  // k链表的后缀 接上当前head
	dummy.Next = kp.Next // k 链表的第一个节点作为新的head
	kp.Next = nil        // 第一个链表的 结尾作为新的tail
	return dummy.Next
}
