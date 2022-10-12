package tree

// 108
func sortedArrayToBST(nums []int) *TreeNode {
	// 找到中间节点 作为根，左边的部分生成一个子树，右边的部分生成一个子树，作为左右节点
	n := len(nums)
	if n == 0 {
		return nil
	}
	var build func(nums []int, start, end int) *TreeNode
	build = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = build(nums, start, mid-1)
		root.Right = build(nums, mid+1, end)
		return root
	}
	return build(nums, 0, n-1)
}
