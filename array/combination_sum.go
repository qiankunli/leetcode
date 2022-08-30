package array

func CombinationSum(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)
}

// 39
func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	doCombinationSum(candidates, 0, path, target, &result)
	return result
}
func doCombinationSum(candidates []int, idx int, path []int, target int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	for i := idx; i < len(candidates); i++ {
		path = append(path, candidates[i])
		doCombinationSum(candidates, i, path, target-candidates[i], result)
		path = path[:len(path)-1]
	}
}
