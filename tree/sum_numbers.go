package tree

import (
	"fmt"
	"math"
)

func SumNumbers(root *TreeNode) int {
	return sumNumbers(root)

}

// 129
// 先遍历所有数字，先存在一起，再求和			这个比较稳妥
func sumNumbers(root *TreeNode) int {
	// 先遍历所有数字，先存在一起，再求和
	// 后序遍历
	if root == nil {
		return 0
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			pathCopy = append(pathCopy, root.Val)
			result = append(result, pathCopy)
			return
		}
		path = append(path, root.Val)
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	s := 0
	for i := 0; i < len(result); i++ {
		tmp := sumPath(result[i])
		fmt.Print(result[i], tmp)
		s += tmp
	}
	return s
}
func sumPath(path []int) int {
	n := len(path)
	s := 0
	for i := n - 1; i >= 0; i-- {
		s += path[i] * int(math.Pow10(n-i-1))
	}
	return s
}

// 后序遍历 这个思路是错的
func sumNumbers2(root *TreeNode) int {

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		lv := 0
		if left >= 0 {
			l := fmt.Sprintf("%d", left)
			nl := len(l)
			lv = root.Val*int(math.Pow10(nl)) + left
		}
		rv := 0
		if right >= 0 {
			r := fmt.Sprintf("%d", right)
			nr := len(r)
			rv = root.Val*int(math.Pow10(nr)) + right
		}
		fmt.Println(left, lv, right, rv)
		if lv+rv > 0 {
			return lv + rv
		}
		return root.Val
	}

	return dfs(root)

}
