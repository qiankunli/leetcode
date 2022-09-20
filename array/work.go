package array

func Max2Profit(cost, profit []int, maxCost int) int {
	return max2Profit(cost, profit, maxCost)
}

// 背包变形，给定n个任务，cost[i] 表示耗时，profit[i] 表示收益，在maxCost 范围内选择2个收益最大的任务
func max2Profit(cost, profit []int, maxCost int) int {
	if len(cost) < 2 {
		return 0
	}
	max := 0
	var dfs func(cost, profit []int, maxCost, idx, depth, total int)
	dfs = func(cost, profit []int, maxCost, idx, depth, total int) {
		if max < total {
			max = total
		}
		n := len(cost)
		if idx >= n {
			return
		}
		if depth >= 2 {
			return
		}

		for i := idx; i < n; i++ {
			c := cost[i]
			if maxCost < c { // 剩余时间不够，剪枝
				continue
			}
			dfs(cost, profit, maxCost-c, i+1, depth+1, total+profit[i])
		}

	}
	dfs(cost, profit, maxCost, 0, 0, 0)
	return max
}
