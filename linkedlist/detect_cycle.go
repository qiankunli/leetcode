package linkedlist

// 142
func detectCycle(head *ListNode) *ListNode {
	// 记录链表经过的节点，如果再一次经过，就是。

	if head == nil {
		return nil
	}
	hash := make(map[*ListNode]bool) // key 是int的话，不能解决val 重复的问题
	p := head
	for p != nil {
		if hash[p] {
			return p
		}
		hash[p] = true
		p = p.Next
	}
	return nil
}
