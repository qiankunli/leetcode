package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

// 02.05
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	curL := l1
	curS := l2
	head := &ListNode{}
	cur := head
	j := 0 // 存储当前的进位
	for curL != nil || curS != nil {
		lv := 0
		if curL != nil {
			lv = curL.Val
		}
		sv := 0
		if curS != nil {
			sv = curS.Val
		}
		v := (lv + sv + j) % 10
		j = (lv + sv + j) / 10
		node := &ListNode{Val: v}
		cur.Next = node

		cur = cur.Next
		if curL != nil {
			curL = curL.Next
		}
		if curS != nil {
			curS = curS.Next
		}
	}
	if j > 0 {
		cur.Next = &ListNode{Val: j}
	}
	return head.Next
}
