package array

// 返回 nums 中第 a 到 第 b 大的数
func FindABLargest(nums []int, a, b int) {
	findABLargest(nums, a, b)
}

// 215
func findABLargest(nums []int, a, b int) {
	if a > b || a >= len(nums) || b < 0 {
		return
	}
	if a < 0 {
		a = 0
	}
	if b >= len(nums) {
		b = len(nums) - 1
	}
	doFindABLargest2(nums, 0, len(nums)-1, a, b)
}
func doFindABLargest(nums []int, start, end, a, b int) {

	if end < start {
		return
	}
	p := nums[start]
	l := start
	r := end
	for l < r {
		for nums[r] >= p && l < r {
			r--
		}
		for nums[l] <= p && l < r {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	// 此时 l ==r 且 [start,l] 都 <= p,    l 天然 >= start
	nums[l], nums[start] = nums[start], nums[l]

	if l > b {
		doFindABLargest(nums, start, l-1, a, b)
		return
	}
	if l < a {
		doFindABLargest(nums, l+1, end, a-l-1, b-l-1)
		return
	}
	doFindABLargest(nums, start, l-1, a, l-1)
	doFindABLargest(nums, l+1, end, l+1, b)
}

// 错误
func doFindABLargest2(nums []int, start, end, a, b int) {

	if end <= start {
		return
	}
	p := (start + end) / 2
	pv := nums[p]
	l := start
	r := end
	for l < r {
		for nums[r] >= pv && l < r {
			r--
		}
		for nums[l] <= pv && l < r {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	// 此时 l ==r  不能保证 [start,l] 都<= pv  因为r 可能会越过 p , lr  在p 左侧进行交换，左侧还是有比p 大的数
	// 且 nums[l] 跟 pv 关系不确定

	if l > b {
		doFindABLargest2(nums, start, l, a, b)
		return
	}
	if l < a {
		doFindABLargest2(nums, l+1, end, a-l-1, b-l-1)
		return
	}
	doFindABLargest2(nums, start, l, a, l)
	doFindABLargest2(nums, l+1, end, l+1, b)
}
