package tree

func LevelOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}

// 102  二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	// 每次读取queue 时， 一次性读取完当前queue的所有元素
	if root == nil {
		return nil
	}
	queue := NewQueue(2000)
	queue.Push(root)
	result := make([][]int, 0)
	for queue.Len() > 0 {
		n := queue.Len()
		nums := make([]int, 0)
		for i := 0; i < n; i++ { // 一次取完
			node := queue.Pop()
			nums = append(nums, node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		result = append(result, nums)
	}

	return result
}
