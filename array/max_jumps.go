package array

import "fmt"

// 1340
func MaxJumps(arr []int, d int) int {
	return maxJumps(arr, d)
}

func maxJumps(arr []int, d int) int {
	// 遍历每一个节点，计算从这个节点 可以跳几步，返回最大值
	max := 0
	//visited := make([]int, len(arr))	// 搜索每深入一层，就跳到了高度更小的位置。因此在搜索的过程中不会出现环
	for i := 0; i < len(arr); i++ {
		c := dfs(arr, d, i)
		fmt.Println(i, c)
		if max < c {
			max = c
		}
	}
	return max
}

// 从idx 往下走可以走几次
func dfs(arr []int, d, idx int) int {

	max := 1 // 即便一个没有，当前节点也算一个
	// 向右
	for i := idx + 1; i <= idx+d; i++ {
		if i < 0 || i >= len(arr) { // 越界
			continue
		}
		if arr[idx] <= arr[i] { // i 后面的不用再遍历了
			break
		}
		c := dfs(arr, d, i)
		if max < c+1 {
			max = c + 1
		}
	}
	// 向左
	for i := idx - 1; i >= idx-d; i-- {
		if i < 0 || i >= len(arr) { // 越界
			continue
		}
		if arr[idx] <= arr[i] { // i 后面的不用再遍历了
			break
		}
		c := dfs(arr, d, i)
		if max < c+1 {
			max = c + 1
		}
	}
	return max
}
