package tree

func PathSum2(root *TreeNode, targetSum int) int {
	return pathSum2(root, targetSum)
}

// 437
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode, targetSum int) int
	dfs = func(root *TreeNode, targetSum int) int {
		if root == nil {
			return 0
		}
		res := 0
		if root.Val == targetSum {
			res++
		}
		l := dfs(root.Left, targetSum-root.Val)
		r := dfs(root.Right, targetSum-root.Val)
		return res + l + r
	}

	croot := dfs(root, targetSum)
	cleft := pathSum2(root.Left, targetSum)
	cright := pathSum2(root.Right, targetSum)

	return croot + cleft + cright
}

// 错误例子
func pathSum3(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == targetSum {
		res++
	}

	// 是两条递归的线路
	l := pathSum3(root.Left, targetSum-root.Val)
	r := pathSum3(root.Right, targetSum-root.Val)

	croot := l + r
	cleft := pathSum2(root.Left, targetSum)
	cright := pathSum2(root.Right, targetSum)

	return croot + cleft + cright
}