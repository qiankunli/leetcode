package tree

// 199

func rightSideView(root *TreeNode) []int {

	// 暴力法，遍历一下，记录每个节点的层数 和 位置，若每层有更靠右的，则覆盖

	// index 为层数，path[i] 为最靠右的节点的location 和值
	path := make([]int, 0)
	var dfs func(root *TreeNode, layer int)
	dfs = func(root *TreeNode, layer int) {
		if root == nil {
			return
		}
		// 将当前节点加入path
		n := len(path)
		if layer >= n {
			path = append(path, root.Val) // 第一次加入
		}
		dfs(root.Right, layer+1) // 总是先看右侧节点
		dfs(root.Left, layer+1)
	}
	dfs(root, 0)
	return path
}

// 一开始是以往有类似的经验，用location记录，但是针对 左子树一直向右 右子树一直像左的场景就不太好处理
func rightSideView2(root *TreeNode) []int {

	// 暴力法，遍历一下，记录每个节点的层数 和 位置，若每层有更靠右的，则覆盖

	// index 为层数，path[i] 为最靠右的节点的location 和值
	path := make([][]int, 0)

	var dfs func(root *TreeNode, layer, location int)
	dfs = func(root *TreeNode, layer, location int) {
		if root == nil {
			return
		}
		// 将当前节点加入path
		n := len(path)
		if layer >= n {
			path = append(path, []int{location, root.Val})
		} else {
			preNode := path[layer]
			if location >= preNode[0] {
				path[layer] = []int{location, root.Val}
			}
		}
		dfs(root.Left, layer+1, location-1)
		dfs(root.Right, layer+1, location+1)
	}
	dfs(root, 0, 0)
	result := make([]int, 0)
	for i := 0; i < len(path); i++ {
		result = append(result, path[i][1])
	}
	return result
}
