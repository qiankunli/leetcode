package graph

func FindJudge(n int, trust [][]int) int {
	return findJudge(n, trust)
}

func findJudge(n int, trust [][]int) int {
	inDegrees := inDegrees(n, trust)
	outDegrees := outDegrees(n, trust)

	// 出度为0的 入度为n-1的，编号从1开始
	for i := 1; i < len(inDegrees); i++ {
		if outDegrees[i] == 0 && inDegrees[i] == n-1 {
			return i
		}
	}
	return -1
}

func inDegrees(n int, trust [][]int) []int {
	result := make([]int, n+1) // 编号从1开始
	for i := 0; i < len(trust); i++ {
		t := trust[i]
		result[t[1]]++
	}
	return result
}
func outDegrees(n int, trust [][]int) []int {
	result := make([]int, n+1) // 编号从1开始
	for i := 0; i < len(trust); i++ {
		t := trust[i]
		result[t[0]]++
	}
	return result
}
