package linkedlist

func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

// 25
// 核心是想到  reverseWithTail 方法，剩下的技术拼凑  reverseWithTail 所需的参数
func reverseKGroup(head *ListNode, k int) *ListNode {

	h := head
	nextH := head
	var newH *ListNode // 记录新数组的头
	for nextH != nil {

		for i := 0; i < k; i++ { // nextH 移动k位，标记下一个要移动group
			nextH = nextH.Next
			if nextH == nil { // 到末尾了
				if newH != nil {
					return newH
				} else {
					return head
				}
			}
		}
		tmpH := reverseWithTail(h, nextH)
		h.Next = nextH // h 此时是最后一个元素
		if newH == nil {
			newH = tmpH
		}
		h = nextH // h 更新下一个group的头节点
	}
	return newH

}

// tail 标记最后一个节点的下一个节点，不参与反序
func reverseWithTail(head, tail *ListNode) *ListNode {

	// 1-> 2-> 3-> 4
	// 1-> nil 2->3->4
	// 2->1->nil 3->4

	// 4->3->2->1
	var pre *ListNode
	cur := head
	for cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseKGroup3(head *ListNode, k int) *ListNode {
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
