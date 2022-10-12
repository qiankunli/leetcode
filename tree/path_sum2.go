package tree

func PathSum2(root *TreeNode, targetSum int) int {
	return pathSum2(root, targetSum)
}

// 437
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 在当前节点必须作为 targetSum 一部分的情况下， 有几个路径
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
// 暴力法，用一个path 跟踪从 根节点到叶子的路径，这样 问题转变成 从一个 array 中找一个连续子序列 和为xx
// 暴力法优化，可以分治下：targetSum 里有根节点，targetSum 里没有根节点。
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