package linkedlist

// 24
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for cur != nil {
		next := cur.Next
		if next == nil {
			break
		}

		// 颠倒cur 和 next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next

		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}
