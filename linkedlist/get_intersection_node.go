package linkedlist

// 160
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	if headA == nil || headB == nil {
		return nil
	}
	cur := headA
	for cur != nil {
		hash[cur] = true
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		if _, ok := hash[cur]; ok {
			return cur
		}
		cur = cur.Next
	}
	return nil
}
