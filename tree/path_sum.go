package tree

func PathSum(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}

// 113
// 34
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

func pathSum4(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil { // 到叶子节点 只能在叶子节点加到path，不然会被加两边
			if target == root.Val {
				pathCopy := make([]int, len(path))
				copy(pathCopy, path)
				pathCopy = append(pathCopy, root.Val)
				result = append(result, pathCopy)
			}
			return
		}

		path = append(path, root.Val)
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return result
}
