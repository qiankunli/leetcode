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
func pathSum5(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, targetSum int) int
	dfs = func(root *TreeNode, targetSum int) int {
		// 这么写的问题在于，如果某个叶子 刚好ok，叶子的左右 都是nil  targetSum == 0 会被执行两次，所以一定要先判断 root == nil
		// 还有一个问题在于，如果是非叶子，并且后面的路径 还存在和为0 的情况  就提前返回了，所以只有到叶子 才是真正的终止条件
		if targetSum == 0 {
			return 1
		}
		if root == nil {
			return 0
		}
		return dfs(root.Left, targetSum-root.Val) +
			dfs(root.Right, targetSum-root.Val)

	}

	// 考虑根
	l := dfs(root.Left, targetSum)

	ll := pathSum5(root.Left, targetSum)
	rr := pathSum5(root.Right, targetSum)

	return l + ll + rr

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
