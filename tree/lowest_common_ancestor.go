package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}

// 236
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	r, _ := doLowestCommonAncestor(root, p, q)
	return r
}

func doLowestCommonAncestor(root, p, q *TreeNode) (*TreeNode, int) {

	count := 0
	// 自己写的根据 slice 生成 tree的工具，-1作为null 值，最后答案可忽略
	if root == nil || root.Val == -1 {
		return root, 0
	}
	if root.Val == p.Val || root.Val == q.Val {
		count = 1
	}
	l, lc := doLowestCommonAncestor(root.Left, p, q)
	if lc == 2 {
		return l, 2
	}
	if count+lc == 2 {
		return root, 2
	}
	r, rc := doLowestCommonAncestor(root.Right, p, q)
	if rc == 2 {
		return r, 2
	}
	if count+lc+rc == 2 {
		return root, 2
	}
	return root, count + lc + rc
}
