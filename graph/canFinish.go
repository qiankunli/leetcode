package graph

// 207

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses) // 第一层课程编号，第二层是 前置课程编号列表
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses) // 第一层课程编号，第二层是 前置课程编号列表
		visited = make([]int, numCourses)
		dfs     func(u int) bool
	)

	// 图的遍历类似于多叉树的遍历，只是要记录下是否遍历过
	dfs = func(u int) bool {

		if visited[u] == 2 { // 已被本轮dfs 访问过
			return true
		}
		if visited[u] == 1 { // 已被本轮dfs 访问过
			return false
		}

		visited[u] = 1
		for _, v := range edges[u] {
			if !dfs(v) {
				return false
			}
		}
		visited[u] = 2
		return true
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}
