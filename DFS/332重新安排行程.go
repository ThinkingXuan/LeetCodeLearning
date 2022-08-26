package main

import "sort"

//给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
//
//
// 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
//
//
//
// 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
//
//
// 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
//
//
//
// 示例 1：
//
//
//输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
//输出：["JFK","MUC","LHR","SFO","SJC"]
//
//
// 示例 2：
//
//
//输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL",
//"SFO"]]
//输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
//解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
//
//
//
//
// 提示：
//
//
// 1 <= tickets.length <= 300
// tickets[i].length == 2
// fromi.length == 3
// toi.length == 3
// fromi 和 toi 由大写英文字母组成
// fromi != toi
//
//
// Related Topics 深度优先搜索 图 欧拉回路 👍 628 👎 0

func findItinerary(tickets [][]string) []string {
	d := map[string][]string{}

	var ans []string

	var dfs func(path string)
	// d: key:飞机出发地点 value 飞机降落的地点集合
	for _, v := range tickets {
		d[v[0]] = append(d[v[0]], v[1])
	}

	// 分机降落的地点集合，按字典排序
	for _, v := range d {
		sort.Strings(v)
	}

	dfs = func(path string) {

		// // 记录结果
		// ans = append(ans, path)

		for len(d[path]) > 0 { // 遍历所有剩余的降落地点
			v := d[path][0]
			d[path] = d[path][1:] //删除已经飞过的降落地点
			dfs(v)
		}
		ans = append([]string{path}, ans...)

	}

	dfs("JFK")
	return ans
}
