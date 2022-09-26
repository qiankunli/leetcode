package tree

func DeleteNode(root *TreeNode, key int) *TreeNode {
	return deleteNode(root, key)
}

// 450
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到，找到了之后，左右孩子如何联系，或者说找谁来替它呢
	// 如果右孩子自己也有孩子呢，找到右子树最小的叶子节点 作为root
	// 如果没有右孩子，就找左子树最大的点来替代它

	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找到 root.Right 最左子节点
		pre := root
		node := root.Right
		for node.Left != nil {
			pre = node
			node = node.Left
		}
		pre.Left = node.Right // 删除与 node 的连接
		node.Left = root.Left
		node.Right = root.Right
		return node
	}
	if key > root.Val {
		right := deleteNode(root.Right, key)
		root.Right = right
		return root
	}
	left := deleteNode(root.Left, key)
	root.Left = left
	return root

}
