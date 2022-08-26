package main

import "math"

//给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
//
//例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
//现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
//
//求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	n := len(routes)

	// 邻接矩阵, 存储站点之间关系
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}

	// res: map[1:[0] 2:[0] 3:[1] 6:[1] 7:[0 1]]

	// rec：
	// false true
	// true false

	// 站点与车之间的关系
	rec := map[int][]int{}

	for i, route := range routes {
		for _, site := range route {
			for _, j := range rec[site] {
				edge[i][j] = true
				edge[j][i] = true
			}
			rec[site] = append(rec[site], i)
		}
	}
	// fmt.Println(rec)

	// for i:=0; i<len(edge); i++ {
	//     for j:=0; j<len(edge[i]); j++ {
	//         fmt.Print(edge[i][j], " ")
	//     }
	//     fmt.Println()
	// }

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}

	var q []int
	for _, bus := range rec[source] {
		dis[bus] = 1
		q = append(q, bus)
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for y, b := range edge[x] {
			if b && dis[y] == -1 {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}

	ans := math.MaxInt32
	for _, bus := range rec[target] {
		if dis[bus] != -1 && dis[bus] < ans {
			ans = dis[bus]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}

	return -1
}
