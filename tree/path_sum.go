package tree

func PathSum(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}

// 113
func pathSum(root *TreeNode, targetSum int) [][]int {

	if root == nil { // 遍历完了
		return nil
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(root *TreeNode, targetSum int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		// 到叶子节点处理一下
		path = append(path, root.Val)
		targetSum -= root.Val

		if root.Left == nil && root.Right == nil && targetSum == 0 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		}
		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum)
		path = path[:len(path)-1]
	}

	dfs(root, targetSum)
	return result
}
