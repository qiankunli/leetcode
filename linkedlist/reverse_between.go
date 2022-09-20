package linkedlist

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	return reverseBetween(head, left, right)
}

// 92
// left 左侧点作为 新的tmpHead，后面的元素头插法插入到 tmpHead 后 知道right
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	tmpHead := &ListNode{Next: head}

	// 找到内部节点的头
	iHead := tmpHead //  内部队列的头
	for i := 0; i < left-1; i++ {
		iHead = iHead.Next
	}
	// 对于 1 -> 2 -> 3 -> 4 -> 5 来说，第一个要移动3 而不是移动2
	// cur 本身不能指向待移动的节点，因为cur 会移动，为此要为cur 记录pre，且要一个变量记录cur.Next以便cur 一会儿指回来。干脆 cur.Next 指向待移动的节点
	cur := iHead.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next

		// 将next 移动到 iHead 之后
		cur.Next = next.Next
		next.Next = iHead.Next
		iHead.Next = next
	}
	return tmpHead.Next
}

