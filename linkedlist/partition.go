package linkedlist

// 86

func partition2(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	small := &ListNode{}
	smallHead := small

	big := &ListNode{}
	bigHead := big

	cur := head
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			big.Next = cur
			big = big.Next
		}
		cur = cur.Next
	}

	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

// 新弄一个小链表，原有链表只保留大于x的节点，问题就在于 维护原有链表的连通性有点麻烦
func partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return nil
	}

	small := &ListNode{}
	smallHead := small

	dummy := &ListNode{Next: head}

	pre := dummy
	cur := head
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next

			pre.Next = cur.Next
		}
		pre = cur
		cur = cur.Next
	}

	next := dummy.Next
	dummy.Next = smallHead.Next
	small.Next = next
	return dummy.Next
}
