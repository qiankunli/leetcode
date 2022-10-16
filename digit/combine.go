package digit

import "fmt"

func Combine(n, k int) [][]int {
	return combine(n, k)
}

// 77
// 给定n,k，返回[1,n] 长为k的子序列
func combine(n, k int) [][]int {
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

func combine2(n int, k int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(idx, k, n int)
	dfs = func(idx, k, n int) {
		if len(path) >= k {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for i := idx; i <= n; i++ {
			path = append(path, i)
			dfs(i+1, k, n)
			path = path[:len(path)-1]
		}

		/*
			这种方式，在dfs(1,k,n) 执行后，立刻会执行 dfs(2,k,n)，path就没有回收的余地了
			path = append(path, idx)
			for i := idx; i <= n; i++ {
				dfs(i+1, k, n)
			}
			path = path[:len(path)-1]
		*/
	}
	dfs(1, k, n)
	return result
}
