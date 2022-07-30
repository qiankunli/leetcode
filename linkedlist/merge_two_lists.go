package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoLists(list1, list2)
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
