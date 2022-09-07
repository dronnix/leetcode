package main

import "fmt"

// The task: https://leetcode.com/problems/unique-paths/
func uniquePaths(m, n int) int {
	mx := make([][]int, m)
	for i := range mx {
		mx[i] = make([]int, n)
		mx[i][0] = 1
	}
	for j := range mx[0] {
		mx[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mx[i][j] = mx[i-1][j] + mx[i][j-1]
		}
	}
	return mx[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 5))
}
