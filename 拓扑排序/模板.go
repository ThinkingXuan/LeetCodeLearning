package main

//有向无环图(Directed Acyclic Graph简称DAG)G进行拓扑排序

/**
numCourses 节点数
prerequisites 先修课程，其中 prerequisites[i] = [ai, bi]， 表示如果要学习课程 ai 则 必须 先学习课程  bi 。
*/

// 判断是否有环
func topSort(numCourses int, prerequisites [][]int) bool {
	// 统计度为0的节点
	indegrees := make([]int, numCourses)
	// 队列
	var queue []int
	// 构建邻接表
	adj := make([][]int, numCourses)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	// 统计度和邻接表
	for i := 0; i < len(prerequisites); i++ {
		cp := prerequisites[i]
		indegrees[cp[0]]++
		adj[cp[1]] = append(adj[cp[1]], cp[0])
	}

	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i) // 度为0的节点入队列
		}
	}

	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]

		numCourses--
		for _, cur := range adj[pre] {
			indegrees[cur]--
			if indegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	return numCourses == 0

}
