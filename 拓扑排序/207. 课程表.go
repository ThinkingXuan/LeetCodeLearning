package main

/*
你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。

例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
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
			queue = append(queue, i) //度为0的节点队列
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
