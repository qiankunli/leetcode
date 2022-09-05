package array

func Search(nums []int, target int) int {
	return search(nums, target)
}

// 33
func search(nums []int, target int) int {
	return kSearch(nums, 0, len(nums)-1, target)
}

// 还是二分查找，针对mid 总有一侧是有序的，另一侧是无序的，则继续切分
func kSearch(nums []int, start, end, target int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	// 检查 mid + 1 和 mid - 1边界场景，因为后面用到了
	if mid+1 >= end {
		if nums[end] == target {
			return end
		}
		return kSearch(nums, 0, mid-1, target)
	}
	if mid-1 <= start {
		if nums[start] == target {
			return start
		}
		return kSearch(nums, mid+1, end, target)
	}
	// 左小右大
	if nums[mid+1] > nums[mid] && nums[mid] > nums[mid-1] {
		// 左侧二分查找
		if index := binarySearch(nums, 0, mid-1, target); index > -1 {
			return index
		}
		// 右侧不确定
		return kSearch(nums, mid+1, end, target)
	}
	// 左小右小 和 左大右大 ，说明mid 刚好在边界，都可以二分查找
	if index := binarySearch(nums, 0, mid-1, target); index > -1 {
		return index
	}
	return binarySearch(nums, mid+1, end, target)
}

func binarySearch(nums []int, start, end, target int) int {
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if target > nums[mid] && end > mid {
		return binarySearch(nums, mid+1, end, target)
	}
	if target < nums[mid] && mid > start {
		return binarySearch(nums, start, mid-1, target)
	}
	return -1
}
