package array

// 540
func SingleNonDuplicate(nums []int) int {
	return singleNonDuplicate(nums)
}

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return doSingleNonDuplicate(nums, n-1)
}
func doSingleNonDuplicate(nums []int, end int) int {
	if end == 0 {
		return nums[end]
	}
	mid := end / 2
	for i := 0; i <= mid; i++ {
		ii := mid + i + 1
		if ii <= end {
			nums[i] = nums[i] ^ nums[ii]
		}
	}
	return doSingleNonDuplicate(nums, mid)
}
