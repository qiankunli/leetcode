package tree

// 二叉树所有没有右兄弟的左叶子节点之和
func leftSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	var dfs func(root *TreeNode, ignore bool)
	dfs = func(root *TreeNode, ignore bool) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && !ignore {
			sum += root.Val
			return
		}

		dfs(root.Left, root.Right == nil)
		dfs(root.Right, true)
	}
	dfs(root, true)
	return sum
}
