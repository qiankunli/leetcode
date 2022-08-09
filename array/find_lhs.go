package array

// 594

func FindLHS(nums []int) int {
	return findLHS(nums)
}
func findLHS(nums []int) int {
	return doFindLHS(nums, 0, len(nums)-1)
}
func doFindLHS(nums []int, left, right int) int {
	if left == right {
		return 0
	}
	v1 := fndLHSWithV(nums, nums[left], left+1, right)
	if v1 > 0 {
		v1 += 1
	}
	v2 := doFindLHS(nums, left+1, right)
	return max(v1, v2)
}

func fndLHSWithV(nums []int, v, left, right int) int {
	// 可以=0 但不能全为0
	count := 0
	equalCount := 0
	for i := left; i <= right; i++ {
		if abs(nums[i]-v) == 1 {
			count++
		}
		if nums[i] == v {
			equalCount++
		}
	}
	if count > 0 {
		return count + equalCount
	}
	return count
}
