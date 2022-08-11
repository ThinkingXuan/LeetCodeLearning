package main

//https://leetcode-cn.com/problems/redundant-connection-ii/submissions/
// 有向图删除
// isTreeAfterRemoveEdge 删一条边之后判断是不是树
func isTreeAfterRemoveEdge(edges [][]int, deleteEdge int) bool {
	initialize()
	for i := 0; i < len(edges); i++ {
		if i == deleteEdge {
			continue
		}
		if same(edges[i][0], edges[i][1]) { // 构成有向环了，一定不是树
			return false
		}
		join(edges[i][0], edges[i][1])
	}
	return true
}

// getRemoveEdge 在有向图里找到删除的那条边，使其变成树
func getRemoveEdge(edges [][]int) []int {
	initialize()
	for i := 0; i < len(edges); i++ { // 遍历所有的边
		if same(edges[i][0], edges[i][1]) { // 构成有向环了，就是要删除的边
			return edges[i]
		}
		join(edges[i][0], edges[i][1])
	}
	return []int{}
}

func findRedundantDirectedConnection(edges [][]int) []int {
	inDegree := make([]int, len(father))
	// 统计各个节点得入度
	for i := 0; i < len(edges); i++ {
		inDegree[edges[i][1]]++
	}

	// 记录入度为2得边
	twoDegree := make([]int, 0)

	// 从后往前判断，因为返回结果需要最后一种情况。
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			twoDegree = append(twoDegree, i)
		}
	}

	// 处理图中情况1 和 情况2
	// 如果有入度为2的节点，那么一定是两条边里删一个，看删哪个可以构成树

	if len(twoDegree) > 0 {
		if isTreeAfterRemoveEdge(edges, twoDegree[0]) {
			return edges[twoDegree[0]]
		}
		return edges[twoDegree[1]]
	}

	// 处理图中情况3
	// 明确没有入度为2的情况，那么一定有有向环，找到构成环的边返回就可以了
	return getRemoveEdge(edges)
}
