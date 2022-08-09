package linkedlist

func HasCycle(head *ListNode) bool {
	return hasCycle(head)
}
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
