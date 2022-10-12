package array

// 26
func removeDuplicates(nums []int) int {
	// 从后往前，发现重复元素，就整体前进一位（复杂度有点高）。统计重复元素的个数c, k = n-c。
	// 从前往后，p 指向当前基准元素，i 指向当前 第一个不与p 相同的元素，将i 移到p+1 的位置，然后改变p 继续移动i。理念就是， 结果里面，后面的元素与前面的都不一样
	// 题解里说，就是将不重复元素移到数组的左侧，形容的精辟
	n := len(nums)
	if n == 1 {
		return 1
	}
	i := 1
	p := 0 // 当前比较的基准
	c := 0 // 重复的个数
	for i < n {

		for i < n && nums[i] == nums[p] {
			i++
			c++
		}
		if i >= n {
			return n - c
		}
		nums[p+1] = nums[i]
		p++
		i++
	}
	return n - c

}
