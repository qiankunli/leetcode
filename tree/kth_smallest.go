package tree

func KthSmallest(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil || k == 0 {
		return -1
	}
	if root.Left == nil && root.Right == nil && k == 1 {
		return root.Val
	}
	lc := Count(root.Left)
	if lc >= k {
		return kthSmallest(root.Left, k)
	} else if lc == k-1 {
		return root.Val
	} else {
		return kthSmallest(root.Right, k-lc-1)
	}
}
