package linkedlist

func IsPalindrome(head *ListNode) bool {
	return isPalindrome(head)
}

// 234
func isPalindrome(head *ListNode) bool {
	mid := Mid(head)
	n := reverse2(mid.Next)
	cur := head
	ncur := n
	for cur != nil && ncur != nil {
		if cur.Val != ncur.Val {
			return false
		}
		if cur == mid {
			break
		}
		cur = cur.Next
		ncur = ncur.Next
	}
	return true
}
