package main

import "fmt"

//编写一个程序，通过填充空格来解决数独问题。
//
//数独的解法需 遵循如下规则：
//
//数字1-9在每一行只能出现一次。
//数字1-9在每一列只能出现一次。
//数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用'.'表示。
//

/*
思路，深度搜索+回溯去做。
每一个空格都有 1-9 九种方法，每次尝试去放其中一个数字，通过check(i,j,num,board)是否合法，不合法直接放另个一个数组，合法就填充这个数组，跳到下一个格积极性
继续放，然后一直递归下去。

出口：可以都到 i==9 j == 9 （代表都是格子都可以放这个数字）

剪枝：题目中已经存在数字的格子直接跳过

*/
func main() {

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j]-'0', " ")
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte) {
	fill(0, 0, board)
}

func fill(i, j int, board [][]byte) bool {
	// 出口
	if j == 9 { // 一行结束
		i++   // 换下一行
		j = 0 // 从第一列开始
		if i == 9 {
			return true
		}
	}

	// 剪枝
	if board[i][j] != '.' {
		return fill(i, j+1, board)
	}

	// 循环填充数字
	for num := byte('1'); num <= byte('9'); num++ {

		if hasConflict(i, j, num, board) {
			continue
		}
		board[i][j] = num
		if fill(i, j+1, board) {
			return true
		}
		board[i][j] = '.'
	}

	return false
}

func hasConflict(row, col int, num byte, board [][]byte) bool {
	// 行和列
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return true
		}
	}

	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+rowStart][j+colStart] == num {
				return true
			}
		}
	}
	return false
}
