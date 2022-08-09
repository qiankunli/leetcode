package array

import "sort"

func TwoSum(nums []int) [][]int {
	return twoSum(nums, 0, 0, len(nums)-1)
}

func twoSum(nums []int, target, start, end int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	i := start
	j := end
	for i < j {
		head := nums[i]
		end := nums[j]
		if head+end == target {
			result = append(result, []int{nums[i], nums[j]})
			i++
			j--
			continue
		}
		if head+end > target {
			j--
			continue
		}
		if head+end < target {
			i++
		}
	}
	return result
}

func twoSum2(nums []int, target, start, end int) [][]int {
	result := make([][]int, 0)
	for i := start; i <= end; i++ {
		for j := i + 1; j <= end; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{nums[i], nums[j]})
			}
		}
	}
	return result
}
