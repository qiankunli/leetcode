package array

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

// 78
func subsets(nums []int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	doSubsets(nums, 0, path, &result)
	return result
}

func doSubsets(nums []int, idx int, path []int, result *[][]int) {
	if idx >= len(nums) {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	// 考虑当前元素
	path = append(path, nums[idx])
	doSubsets(nums, idx+1, path, result)
	// 不考虑当前元素
	path = path[:len(path)-1]
	doSubsets(nums, idx+1, path, result)
}
