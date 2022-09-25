package linkedlist

// 876
func middleNode(head *ListNode) *ListNode {
	p := head
	q := head
	for p != nil && p.Next != nil {
		p = p.Next.Next
		q = q.Next
	}
	return q
}

func middleNode2(head *ListNode) *ListNode {
	p := head
	pre := head
	for p != nil && p.Next != nil {
		p = p.Next.Next // 一次两步
		pre = pre.Next  // 一次一步
	}
	return pre
}
