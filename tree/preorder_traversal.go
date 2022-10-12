package tree

// 144
func preorderTraversal(root *TreeNode) []int {
	path := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return path
}
