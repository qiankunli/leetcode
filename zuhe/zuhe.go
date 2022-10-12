package zuhe

import "fmt"

func Zuhe(n, k int) [][]int {
	return zuhe(n, k)
}

// 给定n,k，返回[1,n] 长为k的子序列
func zuhe(n, k int) [][]int {
	// 边界条件
	if k == 0 {
		return nil
	}
	// 将 n 扩展为数组
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	fmt.Println(nums)
	path := make([]int, 0)
	result := make([][]int, 0)

	var dfs func(nums []int, idx, k int)
	dfs = func(nums []int, idx, k int) {
		if len(path) >= k {
			pathCopy := make([]int, k)
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			// 对于 i 来说，第二层 从i+1 开始的
			dfs(nums, i+1, k) // 面试的时候写成了 idx+1，死活就搞不对了
			path = path[:len(path)-1]
		}
	}

	dfs(nums, 0, k)
	return result
}
