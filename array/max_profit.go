package array

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

// 121
func maxProfit(prices []int) int {
	min := prices[0] // 记录j之前的最小值
	max := 0         // 记录当前的最大利润
	for j := 1; j < len(prices); j++ {
		cur := prices[j] - min
		if max < cur {
			max = cur
		}
		if min > prices[j] {
			min = prices[j]
		}
	}
	return max
}
