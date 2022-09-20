package array

// 202
// 不可能为1 的几个原因：越变越大（经过分析 快乐数 不可能超过xx，不看题解想不出来）；出现了环
// 再进一步优化的话就是查表
func isHappy(n int) bool {
	cur := n
	cache := make(map[int]bool, 0)
	for cur != 1 {
		if _, ok := cache[cur]; ok {
			return false
		}
		cur = happySum(n)
		cache[cur] = true
	}

	return true
}

func happySum(n int) int {
	sum := 0
	for n > 0 {
		if n == 10 {
			sum += 1
			break
		}
		n = n % 10
		sum += n * n
		n = n / 10
	}
	return sum
}
