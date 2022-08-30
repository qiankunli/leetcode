package array

func MajorityElement(nums []int) int {
	return majorityElement(nums)
}

// 169
func majorityElement(nums []int) int {
	m := 0
	c := 0
	for _, v := range nums {
		if v == m {
			c++
		} else if c == 0 {
			m = v
		} else {
			c--
		}
	}
	return m
}
