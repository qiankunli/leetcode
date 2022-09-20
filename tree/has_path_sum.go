package tree

import "fmt"

func HasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSum(root, targetSum)
}

// 112
func hasPathSum(root *TreeNode, targetSum int) bool {
	return sumTree(root, 0, targetSum)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	nums := make([]int, 0)
	have := false
	DFS2(root, nums, func(values []int) bool {
		if sum(values) == targetSum {
			fmt.Println(values)
			have = true
			return true
		}
		return false
	})
	return have
}
func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sumTree(root *TreeNode, curSum, targetSum int) bool {
	if root == nil {
		return false
	}
	reachLeaf := root.Left == nil && root.Right == nil
	if reachLeaf {
		return curSum+root.Val == targetSum
	}
	left := sumTree(root.Left, curSum+root.Val, targetSum)
	right := sumTree(root.Right, curSum+root.Val, targetSum)
	return left || right
}
