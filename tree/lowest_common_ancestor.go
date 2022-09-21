package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}

// 236
// 后序遍历
// 在root 里寻找 p 或q 如果找到则返回。以root 的视角看，如果pq 都在左子树，不管是p 还是q ，先找的的就是 最近公共祖先。
// 如果一个在左 一个在右，根就是root
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil { // 如果左右子树都没找到，说明没有
		return nil
	}
	if left == nil {
		return right // 右子树找到了
	}
	if right == nil {
		return left // 左子树找到了
	}
	return root // 左右都找到了， 公共是根
}

// 之所以麻烦，之所以不好理解，就是没有掌握后序遍历，下面代码的本质是中序遍历
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
