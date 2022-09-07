package array

// 670
// 先拆分，找到最大元素（如果相等，则是靠后的那个）；如果最大元素就是第一个元素，则递归处理后续元素。如果不是， 则交换
func maximumSwap(num int) int {
	nums := split(num)
	doMaximumSwap(&nums, 0)
	return joins(nums)
}
func doMaximumSwap(array *[]int, start int) {
	nums := *array
	if start >= len(nums)-1 {
		return
	}
	max := 0
	max_i := 0
	for i := start; i < len(nums); i++ {
		// 如果相等，就尽量靠后
		if max <= nums[i] {
			max = nums[i]
			max_i = i
		}
	}
	if nums[start] < max {
		nums[start], nums[max_i] = nums[max_i], nums[start]
	} else {
		doMaximumSwap(array, start+1)
	}
}
func split(num int) []int {
	array := make([]int, 0)
	for num > 0 {
		i := num % 10
		array = append(array, i)
		num = num / 10
	}
	nums := make([]int, 0)
	for i := len(array) - 1; i >= 0; i-- {
		nums = append(nums, array[i])
	}
	return nums
}
func joins(nums []int) int {
	sum := 0
	m := 1
	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i] * m
		m *= 10
	}
	return sum
}
