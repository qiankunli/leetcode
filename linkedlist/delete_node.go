package linkedlist

// 237
// 看了题解才知道的，删除node 就必须知道它的pre，既然删不了node，可以删除node.Next 然后让 node.Val = next.Val
func deleteNode(node *ListNode) {

	next := node.Next
	node.Next = next.Next
	// 更新node值
	node.Val = next.Val
}
