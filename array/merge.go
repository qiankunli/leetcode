package array

import (
	"sort"
)

// 56
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	sort.Sort(Range(intervals))
	result := make([][]int, 0)
	i := 1
	p := intervals[0]
	for i < n {
		p1 := intervals[i]
		if p[1] >= p1[0] {
			start := min(p[0], p1[0])
			end := max(p[1], p1[1])
			p = []int{start, end}
			i++
			continue
		}
		result = append(result, p)
		p = p1
		i++
	}
	// 最后一个
	result = append(result, p)
	return result
}

type Range [][]int

func (r Range) Len() int {
	return len(r)
}
func (r Range) Less(i, j int) bool {
	return r[i][0] < r[j][0]
}
func (r Range) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
