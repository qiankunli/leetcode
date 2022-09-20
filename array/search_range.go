package array

// 34
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	if n == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	return search2(nums, 0, n-1, target)
}
func search2(nums []int, start, end, target int) []int {
	l := start
	r := end
	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			return search2(nums, start, mid-1, target)
		} else if target > nums[mid] {
			return search2(nums, mid+1, end, target)
		} else {
			// 难点就是 target 刚好在mid时如何处置
			tr := mid
			for i := mid; i <= end; i++ {
				if target == nums[i] {
					tr = i
				} else {
					break
				}
			}
			tl := mid
			for i := mid; i >= start; i-- {
				if target == nums[i] {
					tl = i
				} else {
					break
				}
			}
			return []int{tl, tr}
		}
	}
	return []int{-1, -1}
}
