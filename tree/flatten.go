package tree

func Flatten(root *TreeNode) {
	flatten(root)
}

// 114

// 本质是递推法 ，先把左左孩子的右子树接上，然后再考虑当前节点的下一个右节点，树的问题什么时候一看dfs 就不好使
func flatten2(root *TreeNode) {
	// 把左子树 放到 父右之间，先移动孩子，再移动父亲
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		// 找到左 最右孩子
		now := cur.Left
		for now.Right != nil {
			now = now.Right
		}
		tmpRight := cur.Right
		cur.Right = cur.Left
		now.Right = tmpRight
		cur.Left = nil

		cur = cur.Right
	}

}

// 先把左右子树变成要求的样子， 再合并左右子树
func flatten(root *TreeNode) {
	// 把左子树 放到 父右之间，先移动孩子，再移动父亲
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	// 返回当前子树的尾，以便连接
	var dfs func(cur *TreeNode) *TreeNode
	dfs = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}

		leftTail := dfs(cur.Left)
		rightTail := dfs(cur.Right)
		// cur 是叶子
		if leftTail == nil && rightTail == nil {
			return cur
		}
		// 左不空，右空
		if leftTail != nil && rightTail == nil {
			cur.Right = cur.Left
			cur.Left = nil
			return leftTail
		}
		// 左空 右不空
		if leftTail == nil && rightTail != nil {
			return rightTail
		}
		leftTail.Right = cur.Right
		cur.Right = cur.Left
		cur.Left = nil
		return rightTail
	}

	dfs(root)
}
