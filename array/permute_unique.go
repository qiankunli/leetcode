package array

// 47

// 因为可以有重复元素，所以一个办法是 记录index 而不是记录索引（最后写入结果时还原为数据）
// 且写入结果时记得去重
func permuteUnique(nums []int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(nums []int)
	dfs = func(nums []int) {
		n := len(nums)
		if len(path) >= n {
			data := make([]int, 0)
			for i := 0; i < len(path); i++ {
				data = append(data, nums[path[i]])
			}
			if !containSlice(result, data) {
				result = append(result, data)
			}
			return
		}
		for i := 0; i < n; i++ {
			if contains(path, i) { // 因为元素可以重复，所以这么判断肯定凉凉
				continue
			}
			path = append(path, i)
			dfs(nums)
			path = path[:len(path)-1]
		}
	}
	dfs(nums)
	return result
}
