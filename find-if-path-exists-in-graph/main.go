package main

import "fmt"

func validPath(n int, edges [][]int, source, destination int) bool {
	if n == 1 {
		return true
	}

	connections := make(map[int][]int, n)

	for _, e := range edges {
		connections[e[0]] = append(connections[e[0]], e[1])
		connections[e[1]] = append(connections[e[1]], e[0])
	}

	return walk(source, destination, connections)
}

func walk(source, destination int, connections map[int][]int) bool {
	nextVertexes, ok := connections[source]
	if !ok {
		return false
	}
	delete(connections, source)

	for _, nextV := range nextVertexes {
		if nextV == destination {
			return true
		}
		if walk(nextV, destination, connections) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {0, 2}, {1, 2}}, 0, 2))
}
