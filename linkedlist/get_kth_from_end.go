package linkedlist

// 22
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	q := head
	c := 0
	for c < k && p != nil {
		p = p.Next
		c++
	}
	if c < k {
		return nil
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}
