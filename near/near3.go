package near

import (
	"math"
	"sort"
)

func Near3(nums []int, target int) int {
	for _, v := range nums {
		if max2 < v {
			max2 = v
		}
	}
	sort.Ints(nums)
	array := split(target)
	return dfs(nums, array, 0, true)
}

var max2 int

// 第一个反应肯定是暴力法 dfs，比如nums 有5个数，每个分叉是5

// 可以优化的点是，每个分叉其实 就三个可能：和当前位相等的，仅小于当前位的，0（也就是当前位不选）。后两种其实是一种情况，一旦小于当前位，后面的位全部以最大值填充即可
func dfs2(nums []int, array []int, idx int, preEqual bool) int {
	if idx >= len(array) { // 终止
		return 0
	}
	if !preEqual {
		cur := max * int(math.Pow10(len(array)-idx-1))
		return cur + dfs(nums, array, idx+1, false)
	}
	for i := len(nums) - 1; i >= 0; i-- { // 如果有小于 array[idx] 的数，实际上最多就执行两次
		if nums[i] > array[idx] {
			continue
		}
		v := dfs(nums, array, idx+1, nums[i] == array[idx])
		if v != -1 {
			cur := nums[i] * int(math.Pow10(len(array)-idx-1))
			return cur + v
		}
	}
	if idx == 0 { // 第一个数特殊处理
		return dfs(nums, array, idx+1, false)
	}
	// preEqual = true ，但是这一轮没有找到 小于 array[idx] 的数，说过这个路径走不通
	return -1
}
