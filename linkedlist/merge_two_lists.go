package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoLists2(list1, list2)
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	tmp, tmp1, tmp2 := h, l1, l2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	// tmp1 和 tmp2 可能没有用完
	for tmp1 != nil {
		tmp.Next = tmp1
		tmp1 = tmp1.Next
		tmp = tmp.Next
	}
	for tmp2 != nil {
		tmp.Next = tmp2
		tmp2 = tmp2.Next
		tmp = tmp.Next
	}
	return h.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	cur1 := list1
	cur2 := list2
	for cur1 != nil || cur2 != nil {
		curV := curV(cur1, cur2)
		newNode := &ListNode{Val: curV.Val}
		cur.Next = newNode
		cur = cur.Next
		if curV == cur1 {
			cur1 = cur1.Next
		} else {
			cur2 = cur2.Next
		}
	}
	return head.Next
}

func curV(node1, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.Val < node2.Val {
		return node1
	}
	return node2
}
