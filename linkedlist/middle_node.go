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
