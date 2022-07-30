package tree

func InvertTree(root *TreeNode) *TreeNode {
	return invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	left := root.Left
	rigth := root.Right
	if left == nil && rigth == nil {
		return root
	}
	root.Left = rigth
	root.Right = left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
