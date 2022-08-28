package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	deps := make([][]int, numCourses)
	for i := range prerequisites {
		deps[prerequisites[i][0]] = append(deps[prerequisites[i][0]], prerequisites[i][1])
	}

	visited := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {

		if !finishable(i, deps, visited) {
			return false
		}
	}

	return true
}

func finishable(start int, deps [][]int, visited []bool) bool {
	if visited[start] {
		return false
	}
	visited[start] = true
	for _, d := range deps[start] {
		if !finishable(d, deps, visited) {
			return false
		}
	}
	visited[start] = false
	deps[start] = nil
	return true
}

func main() {
	deps := [][]int{
		{1, 4},
		{2, 4},
		{3, 1},
		{3, 2},

		{1, 3},
	}
	fmt.Println(canFinish(5, deps))
}
