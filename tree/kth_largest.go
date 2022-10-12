package tree

// 54
func kthLargest(root *TreeNode, k int) int {
	// 左右子树的节点数
	rc := count(root.Right)
	if rc >= k {
		return kthLargest(root.Right, k)
	}
	if rc == k-1 {
		return root.Val
	}
	return kthLargest(root.Left, k-rc-1)

}
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}
