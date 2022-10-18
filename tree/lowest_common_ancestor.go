package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor2(root, p, q)
}

// 236
// 思路是： 分场景，左右子树，都在一个子树上。需要先知道左右子树的情况，才能对特定场景进行处理，所以妥妥的后序遍历。
// 在root 里寻找 p 或q 如果找到则返回。以root 的视角看，如果pq 都在左子树，不管是p 还是q ，先找的的就是 最近公共祖先。
// 如果一个在左 一个在右，根就是root

//  lowestCommonAncestor这个函数不要理解为找公共祖先
// 这里的主要困惑点在于：这样最终能得到 LCA 吗？
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

// 比较顺的思路是，findCount(root,p,q) 寻找root 下 命中pq 的数量，=2 就root 继续下探，=1 =0 表示root 里没有
// 但是findC 比较复杂，本身还要回溯

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

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// pq 都在一边；pq 在两边 此时返回根
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root == p || root == q {
			return root
		}
		l := dfs(root.Left, p, q)
		r := dfs(root.Right, p, q)
		if l != nil && r != nil {
			return root
		}
		if l == nil {
			return r
		}
		return l
	}
	return dfs(root, p, q)
}
