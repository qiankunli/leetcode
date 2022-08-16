package array

func MaxProfit2(prices []int) int {
	return maxProfit2(prices)
}

// 122
func maxProfit2(prices []int) int {
	sum := 0
	buy := prices[0]
	sale := prices[0]
	for j := 1; j < len(prices); j++ {
		if prices[j] >= prices[j-1] {
			sale = prices[j]
			if j < len(prices)-1 {
				continue
			}
		}
		// sale 是局部高点，卖
		if sale > buy {
			sum += sale - buy
		}
		// buy/sale 转移到下一个买卖
		buy = prices[j]
		sale = prices[j]
	}
	return sum
}
