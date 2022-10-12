package tree

import "math"

func IsValidBST(root *TreeNode) bool {
	return isValidBST(root)
}

// 98
// 0405
func isValidBST(root *TreeNode) bool {

	// 左 小于root  右大于root 符合 下探
	if root == nil {
		return true
	}
	// 带区间的递归
	var dfs func(root *TreeNode, min int, max int) bool
	dfs = func(root *TreeNode, min int, max int) bool {
		if root == nil {
			return true
		}
		if root.Val >= max || root.Val <= min {
			return false
		}
		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
