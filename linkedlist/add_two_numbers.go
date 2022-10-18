package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

// 02.05
// 445
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

// 2
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur1 := l1
	cur2 := l2
	dummy := &ListNode{}
	cur := dummy
	extra := 0
	for cur1 != nil || cur2 != nil {
		var v1, v2 int
		if cur1 != nil {
			v1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			v2 = cur2.Val
			cur2 = cur2.Next
		}
		v := v1 + v2 + extra
		if v >= 10 {
			extra = 1
		} else {
			extra = 0
		}
		node := &ListNode{Val: v % 10}
		cur.Next = node
		cur = node
	}
	if extra == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
