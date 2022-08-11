package main

// https://leetcode-cn.com/problems/island-perimeter/
// 对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当这条边为网格的边界或者相邻的另一个格子为水域。 因此，我们可以遍历每个陆地格子，看其四个方向是否为边界或者水域，
// 如果是，将这条边的贡献（即 11）
