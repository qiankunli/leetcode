package array

// 862

// 最终答案是 得到前缀和数组后，遍历前缀和数组搞一点花活
func shortestSubarray(nums []int, k int) int {
	// 滑动窗口 ，其实也是暴力的，只有到最右侧，才知道 有没有凑够k
	// 前缀和

	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = nums[0]
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	min := n + 1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// i j = preSum[j] - preSum[i] + nums[i]
			sum := preSum[j] - preSum[i] + nums[i]
			if sum >= k && min > (j-i+1) {
				min = j - i + 1
			}
		}
	}
	if min == n+1 {
		return -1
	}
	return min
}
