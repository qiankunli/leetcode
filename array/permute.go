package array

func Permute(nums []int) [][]int {
	return permute(nums)
}
func permute(nums []int) [][]int {
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
