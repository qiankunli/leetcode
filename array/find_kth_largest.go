package array

func FindKthLargest(nums []int, k int) int {
	return findKthLargest(nums, k)
}

// 215
func findKthLargest(nums []int, k int) int {

	return doFindKthLargest(nums, 0, len(nums)-1, k)
}

func doFindKthLargest(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}
	left := start
	right := end
	v := nums[left]
	for left < right {
		for nums[right] >= v && left < right {
			right--
		}
		for nums[left] <= v && left < right {
			left++
		}
		swap(nums, left, right)
	}

	swap(nums, start, left)
	// 至此，left 左边比left 小于等于，left 右边比left 大于等于
	if end-left >= k {
		return doFindKthLargest(nums, left+1, end, k)
	} else {
		return doFindKthLargest(nums, start, left, k-(end-left))
	}
}
