package array

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	result := make([][]int, 0)
	i := 0
	for i < len(nums)-2 {
		head := i + 1
		tail := len(nums) - 1
		target := -nums[i]
		for head < tail {
			if nums[head]+nums[tail] == target {
				result = append(result, []int{nums[i], nums[head], nums[tail]})
				head++
				tail--
				// 确保head/tail移动到位
				for head < tail && nums[head] == nums[head-1] {
					head++
				}
				for head < tail && nums[tail] == nums[tail+1] {
					tail--
				}
			}
			if nums[head]+nums[tail] > target {
				tail--
				// 确保head/tail移动到位
				for head < tail && nums[tail] == nums[tail+1] {
					tail--
				}
			}
			if nums[head]+nums[tail] < target {
				head++
				// 确保head/tail移动到位
				for head < tail && nums[head] == nums[head-1] {
					head++
				}
			}

		}
		i++
		// 确保i移动到位
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}

	return result
}
func removeRepeat(nums []int) []int {
	result := make([]int, 0)
	i := 0
	for i < len(nums) {
		result = append(result, nums[i])
		i++
		if i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
	return result
}
func threeSum3(nums []int) [][]int {
	result := make([][]int, 0)
	h := make(map[string]string)
	TraverseChildArrayWithCondition(nums,
		func(nums, bits []int) bool {
			count := 0
			for _, v := range bits {
				if v == 2 {
					count++
				}
			}
			return count <= 3
		},
		func(nums []int) {
			if sum(nums) == 0 && len(nums) == 3 {
				sort.Ints(nums)
				key := join(nums, "_")
				if _, ok := h[key]; !ok {
					result = append(result, nums)
					h[key] = ""
				}
			}
		})
	return result
}

func threeSum2(nums []int) [][]int {
	result := make([][]int, 0)
	h := make(map[string]string)
	TraverseChildArray(nums,
		func(nums []int) {
			if sum(nums) == 0 && len(nums) == 3 {
				sort.Ints(nums)
				key := join(nums, "_")
				if _, ok := h[key]; !ok {
					result = append(result, nums)
					h[key] = ""
				}
			}
		})
	return result
}
