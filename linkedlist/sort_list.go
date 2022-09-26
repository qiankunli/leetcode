package linkedlist

func SortList(head *ListNode) *ListNode {
	return sortList(head)
}

// 148
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// tail 不算的
// sort 干的唯一的活儿就是拆，排序在merge 里做。归并排序的切分方式是取中点
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail { // 相当于到尾部了，只有一个元素
		head.Next = nil
		return head
	}
	// 找到中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	// 排序底层是依次进行的，所以不用担心会断
	return merge(sort(head, mid), sort(mid, tail))
}
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 == nil {
		tmp.Next = tmp2
	}
	if tmp2 == nil {
		tmp.Next = tmp1
	}
	return dummyHead.Next
}
