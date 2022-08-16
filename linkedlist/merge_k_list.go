package linkedlist

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

// 23
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return MergeTwoLists(lists[0], lists[1])
	}
	left := 0
	right := len(lists)
	mid := (left + right) / 2
	leftList := mergeKLists(lists[left:mid])
	rightList := mergeKLists(lists[mid:right])
	return mergeTwoLists(leftList, rightList)
}
