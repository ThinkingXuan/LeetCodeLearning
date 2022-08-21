package main

/*
现在你总共有 numCourses 门课需要选，记为0到numCourses - 1。给你一个数组prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修bi 。

例如，想要学习课程 0 ，你需要先完成课程1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegrees := make([]int, numCourses)

	var queue []int

	adj := make([][]int, numCourses)
	for i := range adj {
		adj[i] = []int{}
	}

	for _, v := range prerequisites {
		indegrees[v[0]]++
		adj[v[1]] = append(adj[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	var res []int

	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		res = append(res, pre)

		numCourses--

		for _, v := range adj[pre] {
			indegrees[v]--
			if indegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 存在环
	if numCourses != 0 {
		return []int{}
	}

	return res
}
