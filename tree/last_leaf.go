package tree

// 返回这个完全二叉树的最后一个节点
func LastLeaf(root *TreeNode) *TreeNode {
	return lastLeaf(root)
}

func lastLeaf(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	hl := getHigh(root.Left)
	hr := getHigh(root.Right)
	if hl == hr { // 左右一样高，说明在右子树
		return lastLeaf(root.Right)
	}
	return lastLeaf(root.Left)
}
func getHigh(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	hl := getHigh(root.Left)
	hr := getHigh(root.Right)
	max := hl
	if max < hr {
		max = hr
	}
	return max + 1
}
