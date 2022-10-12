package tree

// 010
func isSymmetric(root *TreeNode) bool {
	// dfs 搞不定，层次遍历 每一层的元素是否是对称的
	// 看了题解，真尼玛绝

	if root == nil {
		return true
	}
	var dfs func(left, right *TreeNode) bool
	// 绝就绝在不像往常一样，透传的是根
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		}
		// 都不为null
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)

}
