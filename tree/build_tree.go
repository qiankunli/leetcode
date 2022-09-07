package tree

// 105
func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	return doBuildTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

// 分析例子，从先序可以拿到根节点，对于中序来说，根节点左边的是左子树，右边是右子树，这样就可以把问题递归化了
func doBuildTree(preorder []int, inorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
	if pStart > pEnd {
		return nil
	}
	if pEnd == pStart {
		return &TreeNode{Val: preorder[pStart]}
	}
	rootVal := preorder[pStart]
	root := &TreeNode{Val: rootVal}
	for i := iStart; i <= iEnd; i++ {
		if inorder[i] == rootVal {
			len := i - iStart
			root.Left = doBuildTree(preorder, inorder, pStart+1, pStart+len, iStart, i-1)
			root.Right = doBuildTree(preorder, inorder, pStart+len+1, pEnd, i+1, iEnd)
		}
	}
	return root
}
