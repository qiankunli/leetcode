package array

// 217
func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			hash[v] = hash[v] + 1
			if hash[v] >= 2 {
				return true
			}
		} else {
			hash[v] = 1
		}
	}
	return false
}
