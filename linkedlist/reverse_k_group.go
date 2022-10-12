package linkedlist

func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

// 25

// 没事手撸一下
func reverseKGroup4(head *ListNode, k int) *ListNode {
	// 首先得有一个基本函数 reverse(head,tail), 负责反转[head,tail)，且确保反转后list 没断，一次反转k个，反转完再弄下一个批次

	dummy := &ListNode{Next: head}
	pre := dummy // 指向子链前一个节点
	h := head    // 指向子链第一个节点
	t := head    // 指向子链第一个节点，移动后指向最后一个节点的下一个节点

	for {
		// t 后移 k个位置
		i := 0
		for i < k {
			if t == nil { // 碰到了尾部
				break
			}
			t = t.Next
			i++
		}
		if i < k { // 长度不够，提前结束
			break
		}
		newH := reverseTail(h, t) // h 指向t， 返回新的 head
		pre.Next = newH           // 串联反转后的子链
		pre = h                   // 指向下一个k 子链的前一个节点
		next := h.Next
		t = next // 指向子链第一个节点，移动后指向最后一个节点的下一个节点
		h = next

	}
	return dummy.Next
}

// 反转head 和 tail（不包含tail），返回最新的head，且新的尾节点可以连接tail
func reverseTail(head, tail *ListNode) *ListNode {
	pre := tail
	cur := head
	for cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

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
