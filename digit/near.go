package digit

import (
	"fmt"
	"math"
	"sort"
)

var maxNearMax int

// 寻找小于target 的最大数
func Near(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if target == 0 {
		return 0
	}
	sort.Ints(nums)
	targetNums := split(target)
	return near2(nums, targetNums, true, 0)
	//return maxNearMax
}

// 还是直接暴力回溯法比较简洁
func near(nums []int, target int, num int) {

	if calLens(num) >= calLens(target) {
		return
	}
	for i := 0; i < len(nums); i++ {
		num = num*10 + nums[i] // 考虑 nums[i], 缺点就在于 num 一开始位数不够，肯定会小于 target，会出现 99 小于 1000 的情况，但9 因为大于1 不能作为最高位的选择
		if maxNearMax < num && num < target {
			maxNearMax = num
		}
		near(nums, target, num)
		num = num / 10 // 恢复信息
	}
}

// 这靠面试现场肯定搞不定
// 从暴力回溯出发，优化点1：仅选择比当前位 等于或小于的数字，但这个优化点要注意，如果找不到，当前位不选，后面的就直接由最大的数字填充了
// preEq 前面是否都是相等的，index 当前遍历到哪了
func near2(nums []int, targetNums []int, preEq bool, index int) int {
	if index >= len(targetNums) {
		return 0
	}
	if preEq {
		// 从nums 中找到数字处理
		for i := len(nums) - 1; i >= 0; i-- {
			// 此处先暂时忽略 最后一位是否相等的处理
			// 要么就是 当前层 找不到 小于 targetNums[index] 的值，要么就是 孩子节点找不到
			if nums[i] <= targetNums[index] {
				temp := near2(nums, targetNums, nums[i] == targetNums[index], index+1)
				if temp == -1 {
					continue
				}
				return nums[len(nums)-1]*int(math.Pow10(len(targetNums)-index-1)) + temp
			}
		}
		// 从nums 中没找到数字
		// ...
		if index == 0 { // 第一个数字无法满足
			return near2(nums, targetNums, false, index+1)
		}
		// 前面都一样，但从nums 中没找到数字可用，说明此路不通，回退。第一层除外。
		return -1
	}
	return nums[len(nums)-1]*int(math.Pow10(len(targetNums)-index-1)) + near2(nums, targetNums, false, index+1)
}

// 字节一面写的回答，已经乱掉了
//func near2(nums []int, target int,path *[]int) int {
//	if target == 0 {
//		return 0
//	}
//	targetLen := calLens(target)
//	curNumber := target / int(math.Pow10(targetLen-1))
//	curNearMax := findNearMax(nums, curNumber)
//	if curNearMax == curNumber {
//		path = append(path, curNumber)
//		re := near(nums, target-curNumber*int(math.Pow10(targetLen-1)))
//		if re == 0 {
//			return 0
//		} else {
//			near(nums, target-curNumber*int(math.Pow10(targetLen-1))-int(math.Pow10(targetLen-2)))
//		}
//	}
//	max := findMax(nums)
//	if curNearMax == -1 {
//		sum := 0
//		for i := 0; i < targetLen-1; i++ {
//			sum = sum * 10 + max
//		}
//		return sum
//	}
//	if curNearMax < curNumber {
//		for i := 0; i < targetLen; i++ {
//			path = append(path, max)
//		}
//	}
//	return 0
//}
func calLens(num int) int {
	s := fmt.Sprintf("%d", num)
	return len(s)
}
func findNearMax(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			result = nums[i]
		}
	}
	return result
}
func findNearMax2(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			result = nums[i]
		}
	}
	return result
}
func findMax(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
func findMin(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}
func joinNums(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i] * int(math.Pow10(n-i-1))
	}
	return sum
}
func countNot(nums []int, target int) int {
	c := 0
	for _, v := range nums {
		if v != target {
			c++
		}
	}
	return c
}

func findMaxNumber(target int) int {
	max := 0
	for target > 0 {
		cur := target % 10
		if max < cur {
			max = cur
		}
		target = target / 10
	}
	return max
}

func split(target int) []int {
	nums := make([]int, 0)
	for target > 0 {
		cur := target % 10
		nums = append(nums, cur)
		target = target / 10
	}
	result := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
	}
	return result
}
