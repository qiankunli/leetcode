package linkedlist

import "fmt"

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	return removeNthFromEnd2(head, n)
}

// 19
// 反转；删除正数第n个；反转    有点想复杂了。
// 可以直接双指针的

func removeNthFromEnd2(head *ListNode, n int) *ListNode {

	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	// second 最后要停留在最n个数的前一个，所以first 和 second 要相差n个
	second.Next = second.Next.Next
	return dummy.Next
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := reverse(head)

	if n == 1 {
		return reverse(h.Next)
	}

	c := 1
	var pre *ListNode
	cur := h
	for c < n && cur != nil {
		pre = cur
		cur = cur.Next
		c++
	}
	// n 大于 链表长度
	if c < n {
		return nil
	}
	// cur 指向待删除的节点，删除它
	pre.Next = cur.Next
	return reverse(h)
}

// 1 -> 2 -> 3 -> 4
// 1-> nil 2->3->4
// 2->1->nil 3->4
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func print(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
