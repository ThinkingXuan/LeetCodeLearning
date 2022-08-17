package main

// 全局变量
var (
	n      = 1005 // 节点数量3 到 1000
	father = make([]int, n)
	size   = make([]int, n) // 每个连通分量的个数

)

// 并查集初始化
func initialize() {
	for i := 0; i < n; i++ {
		father[i] = i
		size[i] = 1
	}
}

// 并查集里寻根的过程
func find(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

// 将v->u 这条边加入并查集
func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
	size[u] += size[v]
}

// 判断 u 和 v是否找到同一个根，
func same(u, v int) bool {
	u = find(u)
	v = find(v)
	return u == v
}

// 统计连通分量(求子树数量)
// 出入节点个数
func connectSize(n int) int {
	res := 0
	for i, v := range father[:n] {
		if i == v {
			res++
		}
	}
	return res
}

// 统计连通分量中最大的数量
func connectMaxSize(n int) int {
	maxSize := 0
	for i, v := range father[:n] {
		if i == v {
			maxSize = max(maxSize, size[i])
		}
	}
	return maxSize
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 冗余连接问题，（删除图中得一条边，然后构成树。。返回要删除得边）
// https://leetcode-cn.com/problems/redundant-connection/
func findRedundantConnection(edges [][]int) []int {
	initialize()
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		if same(a, b) {
			return edges[i]
		} else {
			join(a, b)
		}
	}
	return []int{}
}
