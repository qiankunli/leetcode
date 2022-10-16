package array

// 27
func removeElement(nums []int, val int) int {
	// 两个指针，q一个指向碰到的 val ，p另一个指向不等于 val的元素，然后复制，但真正操作起来很复杂。当着想着不等于val 的元素不动
	// 这个思路也是两个指针，但q 是依次连续右移的，不等于val 的元素整体向前搬迁

	n := len(nums)
	if n == 0 {
		return 0
	}

	p := 0
	q := 0
	c := 0
	for p < n {
		for p < n && nums[p] != val { // p 找到不是q 的元素，并移动
			c++
			nums[q] = nums[p]
			q++
			p++
		}

		p++
	}
	return c
}
func removeElement2(nums []int, val int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	p := 0
	q := 0
	c := 0
	for p < n {
		for p < n && nums[p] == val {
			p++
		}
		if p == n {
			break
		}
		c++
		nums[q] = nums[p]
		q++
		p++
	}
	return c
}
