package array

import (
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	return combinationSum2(candidates, target)
}

// 40
func combinationSum2(candidates []int, target int) [][]int {
	path := make([]int, 0) // 存下标
	result := make([][]int, 0)
	sort.Ints(candidates) // 使得相同的元素放在一起
	doCombinationSum2(candidates, 0, path, target, &result)
	return result
}
func doCombinationSum2(candidates []int, idx int, path []int, target int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		pathCopy := make([]int, 0)
		for _, index := range path {
			pathCopy = append(pathCopy, candidates[index])
		}
		*result = append(*result, pathCopy)
		return
	}
	for i := idx; i < len(candidates); i++ {
		candidate := candidates[i]
		// 如果当前值和前一个值一样，则在遍历前一个值的时候，已经把当前子树遍历过了，不用继续遍历了
		if i > 0 && candidate == candidates[i-1] {
			// 如果前一个值也在路径里，则说明 target 需要重复
			if !(len(path) > 0 && path[len(path)-1] == i-1) {
				continue
			}
		}
		path = append(path, i)
		doCombinationSum2(candidates, i+1, path, target-candidate, result)
		path = path[:len(path)-1]
	}
}
