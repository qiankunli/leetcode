package array

// 88
func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[m+i] = nums2[i]
		}
		return
	}
	i1 := 0 // 指示nums1 待比较的元素
	i2 := 0
	curN := m // 每次插入后的尾部
	for i1 < len(nums1) && i2 < n {
		if nums1[i1] > nums2[i2] {
			curN++
			insert(nums1, curN, i1, nums2[i2])
			i2++
			i1++
		} else {
			i1++
		}
	}
	for i2 < n {
		nums1[curN] = nums2[i2]
		curN++
		i2++
	}
}
func insert(nums []int, n, index, v int) {
	for i := n - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = v
}
