package array

// 88

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge3(nums1, m, nums2, n)
}

// 从后到前开始判断
func merge3(nums1 []int, m int, nums2 []int, n int) {
	// 三个指针   nums2 最新待决定  nums1 最新待决定，  nums1 最大值

	// 边界判断
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	// p1 和 p2 可能没有指到底
	//for p1 >= 0 {		// nums1 本来就是，可以不动
	//	nums1[p] = nums1[p1]
	//	p1--
	//	p--
	//}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

// 第一反应   nums2 从头到尾插入到 nums1 ，nums1 部分元素整体后移
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
