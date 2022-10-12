package tree

// 543
// 第一个反应是 左子树的深度 + 右子树的深度 + 1
// 之后很快想到了反例，比如left 左右子树深度很大， right 深度很小，则最后答案出在 left 为根的子树上
// 之后想到这是一个从下到上的过程，从下到上依次求每个 节点的直径，用max 记录最大值，后序遍历
// 结合一下，我们需要一个函数 求 left right 的最大深度，以便求root 的直径。拿到root 的直径后，又要更新max 更新最大直径

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if m < left+right { //  [4,2,1,3]  的长度为3
			m = left + right
		}
		return max(left, right) + 1
	}
	dfs(root)
	return m
}
