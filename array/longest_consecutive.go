package array

// 128

// 滑动窗口不行
// 类似，最长递增子序列。dp  一维dp 是不行的
// 如果用回溯，子序列，最大值 - 最小值 = 数组长度 且都不一样
// 看题解发现想多了
// 超时
func longestConsecutive(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	hash := make(map[int]bool)
	for i := 0; i < n; i++ {
		hash[nums[i]] = true
	}
	max := 1
	for k, _ := range hash {
		num := k
		numLen := 1
		for hash[num+1] {
			num++
			numLen++
			if max < numLen {
				max = numLen
			}
		}
	}
	return max
}
