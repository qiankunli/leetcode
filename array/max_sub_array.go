package array

import "math"

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}
func maxSubArray(nums []int) int {
	return doMaxSubArray(nums, 0, len(nums)-1)
}
func doMaxSubArray(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	if end-start == 1 {
		return max3(nums[start], nums[end], nums[start]+nums[end])
	}
	mid := (start + end) / 2
	left := doMaxSubArray(nums, start, mid-1)
	right := doMaxSubArray(nums, mid+1, end)
	cross := doMaxSubArrayWithMid(nums, mid, start, end)
	return max3(left, right, cross)
}
func doMaxSubArrayWithMid(nums []int, mid, start, end int) int {
	sum1 := 0
	max1 := math.MinInt
	for i := mid - 1; i >= start; i-- {
		sum1 += nums[i]
		if max1 < sum1 {
			max1 = sum1
		}
	}
	sum2 := 0
	max2 := math.MinInt
	for i := mid + 1; i <= end; i++ {
		sum2 += nums[i]
		if max2 < sum2 {
			max2 = sum2
		}
	}
	return nums[mid] + max(max1, 0) + max(max2, 0)
}

// 递归不过是遍历的另一种说法
// 遍历连续字串的迭代方式
func maxSubArray2(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
