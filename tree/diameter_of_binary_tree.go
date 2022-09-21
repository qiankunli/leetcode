package tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if m < left+right {
			m = left + right
		}
		return max(left, right) + 1
	}
	dfs(root)
	return m
}
