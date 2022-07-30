package split

import "math"

func Split() []int {
	total := 100
	// 硬性要求：必须大于guarantee 小于 capacity
	// 软性要求：必须尽可能贴和weight 和 request
	// 如果request = 0, 则weight 视为 0，表示不参与分配
	// 如果request = 0, 则capacity = guarantee
	weights := []int{0, 20, 30, 40} // weight * total 后的值
	guarantees := []int{30, 0, 10, 0}
	capacities := []int{30, 60, 70, 60} // real capacity
	requests := []int{0, 60, 30, 10}

	min_cost := math.MaxInt32
	var result []int
	for r1 := guarantees[0]; r1 <= capacities[0]; r1++ {
		for r2 := guarantees[1]; r2 <= capacities[1]; r2++ {
			for r3 := guarantees[2]; r3 <= capacities[2]; r3++ {
				for r4 := guarantees[3]; r4 <= capacities[3]; r4++ {
					if (r1 + r2 + r3 + r4) != total {
						continue
					}
					costw := EulerDistance([]int{r1, r2, r3, r4}, weights)
					costr := EulerDistance([]int{r1, r2, r3, r4}, requests)
					cost := 2*costr + costw
					if min_cost > cost {
						min_cost = cost
						result = []int{r1, r2, r3, r4}
					}
				}
			}
		}
	}
	return result
}

func EulerDistance(nodeX, nodeY []int) int {
	sum := 0
	for i := 0; i < len(nodeX); i++ {
		sum += square(nodeX[i] - nodeY[i])
	}
	return sum
}
func square(x int) int {
	return x * x
}
