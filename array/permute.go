package array

func Permute(nums []int) [][]int {
	return permute(nums)
}

// 46 全排列

func permute(nums []int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(nums []int, idx int)
	dfs = func(nums []int, idx int) {
		n := len(nums)
		if idx >= n {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}
		for i := 0; i < n; i++ {
			if pathContains(path, nums[i]) {
				continue
			}
			path = append(path, nums[i])
			dfs(nums, idx+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return result
}
func pathContains(path []int, target int) bool {
	for _, v := range path {
		if v == target {
			return true
		}
	}
	return false
}

func permute2(nums []int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	doPermute(nums, 0, path, &result)
	return result
}

func doPermute(nums []int, idx int, path []int, result *[][]int) {
	if len(path) == len(nums) {
		pathCopy := make([]int, len(path))
		for i, v := range path {
			pathCopy[i] = nums[v]
		}
		*result = append(*result, pathCopy)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(path, i) {
			continue
		}
		path = append(path, i)
		doPermute(nums, i, path, result)
		path = path[:len(path)-1]
	}
}
