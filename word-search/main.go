package main

import "fmt"

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}

	for row := range board {
		for col := range board[row] {
			if board[row][col] == word[0] && dfs(row, col, word, board, used) {
				return true
			}
		}
	}
	return false
}

func dfs(row, col int, word string, board [][]byte, used [][]bool) bool {
	if word == "" {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if used[row][col] {
		return false
	}
	if board[row][col] != word[0] {
		return false
	}
	used[row][col] = true
	if dfs(row+1, col, word[1:], board, used) || dfs(row-1, col, word[1:], board, used) ||
		dfs(row, col+1, word[1:], board, used) || dfs(row, col-1, word[1:], board, used) {
		return true
	}
	used[row][col] = false
	return false
}

func main() {
	b := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(b, "SEE"))

}
