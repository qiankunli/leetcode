package linkedlist

// 82
func deleteDuplicates(head *ListNode) *ListNode {
	// dummy 节点，两个指针，pre p，两层循环，内层找到第一个跟p 值不同的节点

	dummy := &ListNode{Next: head}
	p := head
	pre := dummy // 作为p 前面的节点
	for p != nil {
		cur := p.Next
		// 往下找，直到找到与p 值不同的节点
		for cur != nil {
			if cur.Val != p.Val {
				break
			}
			cur = cur.Next
		}
		if cur == p.Next { // p 后实际上没重复
			pre = p
			p = p.Next
		} else { // p后有重复
			pre.Next = cur
			p = cur
		}
	}
	return dummy.Next
}
