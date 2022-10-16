package digit

func NextPermutation(nums []int) {
	nextPermutation(nums)
}

// 31

func nextPermutation2(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	// 从右向左找nums[p] 使得nums[p] < nums[p+1]
	p := n - 2
	for p >= 0 && nums[p] <= nums[p+1] {
		p--
	}
	if p < 0 { // 说明nums 是降序
		// 实际的都是降序，此时用快排没必要
		up(nums, 0, n-1)
		return
	}
	// 从右向左找nums[q] 使得nums[q] > nums[p] 且 nums[q] 是 p+1,n 最小的
	q := n - 1
	for q > p && nums[q] <= nums[p] { // 这种方式不保证 nums[q] 是 p+1,n 最小的
		q--
	}
	// 交换pq
	nums[p], nums[q] = nums[q], nums[p]
	up(nums, p+1, n-1) // 此处可能乱序
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	// 从右向左找nums[p] 使得nums[p] < numps[p+1]
	p := n
	for i := n - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			p = i - 1
			break
		}

	}
	if p == n { // 说明nums 是降序
		// 实际的都是降序，此时用快排没必要
		up(nums, 0, n-1)
		return
	}
	// 从右向左找nums[q] 使得nums[q] > numps[p] 且 nums[q] 是 p+1,n 最小的
	q := n
	min := 101
	for i := n - 1; i > p; i-- {
		if nums[i] > nums[p] {
			if min > nums[i] {
				min = nums[i]
				q = i
			}
		}
	}
	// 交换pq
	nums[p], nums[q] = nums[q], nums[p]
	up(nums, p+1, n-1) // 此处可能乱序
}

// 仅针对降序的排序
func up2(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	for i := start; i <= mid; i++ {
		nums[i], nums[end-i] = nums[end-i], nums[i]
	}
}
func up(nums []int, start, end int) {
	if start >= end {
		return
	}
	p := nums[start]
	//l := start + 1 因为最后有一个  l 与 start 交换，这回会导致一定会交换一次
	l := start
	r := end
	for l < r {
		for nums[r] >= p && l < r { // l是不能越过r 的，从算法思想上也没必要
			r--
		}
		for nums[l] <= p && l < r {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[start] = nums[start], nums[l]
	up(nums, start, l-1)
	up(nums, l+1, end)

}
