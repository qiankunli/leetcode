package tree

// 删除不在[a,b] 范围内的节点
func TrimBST(root *TreeNode, a, b int) *TreeNode {
	return trimBST(root, a, b)
}

// 669
func trimBST(root *TreeNode, a, b int) *TreeNode {
	if root == nil {
		return nil
	}
	// 根 + 右子树超出界限，删掉
	if root.Val > b {
		deleteTreeNode(root.Right)
		return trimBST(root.Left, a, b) // 根节点转移为 左节点，继续尝试
	}
	// 根 + 左子树超出界限，删掉
	if root.Val < a {
		deleteTreeNode(root.Left)
		return trimBST(root.Right, a, b) // 根节点转移为 右节点，继续尝试
	}
	// 根节点在 [a,b] 内，根节点不会转移
	left := trimBST(root.Left, a, b)
	right := trimBST(root.Right, a, b)
	root.Left = left   // 重新建立与删掉后节点的连接
	root.Right = right // 重新建立与删掉后节点的连接
	return root
}
func deleteTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	deleteTreeNode(root.Left)
	deleteTreeNode(root.Right)
	root.Left = nil
	root.Right = nil
}
