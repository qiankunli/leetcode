package array

import "math"

func MaxProduct(nums []int) int {
	return maxProduct(nums)
}

// 152
// 乘积最大子数组
func maxProduct(nums []int) int {
	// 不能排序
	// 滑动窗口  碰到正数就右移， 碰到负数且负数的下一个是正数，窗口就结束？ 不一定， 所以滑动窗口解决不了该问题
	// 那就只能暴力法，不过用dp 来加速。公式不看答案真的推导不出来

	n := len(nums)
	if n == 0 {
		return 0
	}
	maxn := make([]int, n)
	maxn[0] = nums[0]
	minn := make([]int, n)
	minn[0] = nums[0]
	// max[i] = max(max[i-1]*nums[i],nums[i],min[nums-i]*nums[i])
	// min[i] = min(min[i-1]*nums[i],nums[i],max[nums-i]*nums[i])
	for i := 1; i < n; i++ {
		maxn[i] = max3(maxn[i-1]*nums[i], nums[i], minn[i-1]*nums[i])
		minn[i] = min3(maxn[i-1]*nums[i], nums[i], minn[i-1]*nums[i])
	}
	// 返回maxn 里的最大值
	r := math.MinInt
	for i := 0; i < n; i++ {
		if r < maxn[i] {
			r = maxn[i]
		}
	}
	return r
}
