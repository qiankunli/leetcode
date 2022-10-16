package array

// 35
func searchInsert(nums []int, target int) int {
	// logn 就预示了使用二分法
	n := len(nums)
	if n == 0 {
		return 0
	}

	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {

			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l

}
