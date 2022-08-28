package main

import "fmt"

// Task: https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
	}

	num := 0
	for i := range grid {
		for j := range grid[i] {
			if visited[i][j] || grid[i][j] == '0' {
				continue
			}
			discoverIsland(i, j, grid, visited)
			num++
		}
	}
	return num
}

func discoverIsland(r, c int, grid [][]byte, visited [][]bool) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] == '0' || visited[r][c] {
		return
	}
	visited[r][c] = true
	discoverIsland(r+1, c, grid, visited)
	discoverIsland(r-1, c, grid, visited)
	discoverIsland(r, c+1, grid, visited)
	discoverIsland(r, c-1, grid, visited)
}

func main() {
	g := [][]byte{
		{'0', '0', '0', '0'},
		{'0', '1', '1', '0'},
		{'0', '0', '1', '0'},
		{'0', '0', '0', '0'},
	}
	fmt.Println(numIslands(g))

}
