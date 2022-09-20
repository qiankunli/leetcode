package tree

import "fmt"

// 662
func WidthOfBinaryTree(root *TreeNode) int {
	return widthOfBinaryTree(root)
}

// 遍历一遍，记录一个map[int]int  key=深度，value=节点个数，最后给出最大值
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := make([]int, 0)  // 记录左子树最靠左的 偏移
	right := make([]int, 0) // 记录右子树最靠右的 偏移
	var dfs func(root *TreeNode, depth, width int)
	dfs = func(root *TreeNode, depth, width int) {
		if root == nil {
			return
		}
		if width < 0 {
			if depth >= len(left) {
				left = append(left, width)
			} else {
				left[depth] = width
			}

		}
		if width > 0 {
			if depth >= len(right) {
				right = append(right, width)
			} else {
				right[depth] = width
			}
		}
		dfs(root.Left, depth+1, width-1)
		dfs(root.Right, depth+1, width+1)
	}
	dfs(root, 0, 0)
	fmt.Println(left, right)
	nl := len(left)
	nr := len(right)
	if nl == 0 || nr == 0 { // 某个子树没有
		return 0
	}
	l := nl
	if l > nr {
		l = nr
	}
	fmt.Println(l)
	return right[l-1] - left[l-1]
}
