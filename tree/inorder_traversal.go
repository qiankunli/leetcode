package tree

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	if root.Left != nil {
		left := inorderTraversal(root.Left)
		for _, v := range left {
			result = append(result, v)
		}
	}
	result = append(result, root.Val)

	if root.Right != nil {
		right := inorderTraversal(root.Right)
		for _, v := range right {
			result = append(result, v)
		}
	}
	return result
}
