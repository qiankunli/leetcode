package tree

import "math"

func MaxPathSum(root *TreeNode) int {
	return maxPathSum(root)
}

// 51
func maxPathSum(root *TreeNode) int {
	max := math.MinInt
	// 这道题必须得从底向上考虑，不能从上向下考虑，这要求额外拎出一个变量来记录 从下到上的最大值，传统的递归做不到
	var doMaxPathSum func(root *TreeNode) int
	doMaxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max2(doMaxPathSum(root.Left), 0)
		right := max2(doMaxPathSum(root.Right), 0)
		//max := max3(left, right, all)	// 在这个地方想当然了，花了很长时间

		// 最大路径和 和 最大贡献值是两个东西，这是理解本题的关键，最大路径和可以是 根 + 左 + 右，最大贡献值只能是 根 + 左或 根 + 右
		// 最大贡献值
		val := max3(root.Val, root.Val+left, root.Val+right)
		// 最大路径和
		max = max2(max, root.Val+left+right)
		return val
	}
	doMaxPathSum(root)
	return max
}

func max2(l, r int) int {
	if l > r {
		return l
	}
	return r
}
