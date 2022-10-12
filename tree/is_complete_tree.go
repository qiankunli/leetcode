package tree

// 958

func IsCompleteTree(root *TreeNode) bool {
	return isCompleteTree(root)
}

// 如果左节点没有，但是有右节点，就说明不是
// 左节点个数小于右节点的个数
// 如果右节点有数据，左节点就必须满。自己想到了层次遍历，同层次节点的遍历，只有bfs 才能创造这样的机会
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := NewQueue(101)
	queue.Push(root)
	for queue.Len() > 0 {
		n := queue.Len()
		flag := false
		// 我写的时候 一直在想办法判断  如果next 为空、有没有孩子，则cur 什么情况算非法，细节就多了。
		for i := 0; i < n; i++ {
			node := queue.Pop()
			if node != nil && flag { // 如果有node 不为nil 但之前碰到了nil 则返回false
				return false
			} else {
				flag = true
			}
			queue.Push(root.Left)
			queue.Push(root.Right)
		}
	}
	return true
}
