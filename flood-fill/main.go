package main

import "fmt"

// The task: https://leetcode.com/problems/flood-fill/

func floodFill(image [][]int, sr, sc, color int) [][]int {
	origc := image[sr][sc]
	if origc == color {
		return image
	}
	walk(image, sr, sc, color, origc)
	return image
}

func walk(image [][]int, sr, sc, color, origc int) {
	if sr >= len(image) || sr < 0 {
		return
	}
	if sc >= len(image[0]) || sc < 0 {
		return
	}

	if image[sr][sc] != origc {
		return
	}
	image[sr][sc] = color

	walk(image, sr+1, sc, color, origc)
	walk(image, sr-1, sc, color, origc)
	walk(image, sr, sc+1, color, origc)
	walk(image, sr, sc-1, color, origc)
}

func main() {
	fmt.Println(floodFill([][]int{{0, 1}, {0, 0}}, 0, 0, 2))
}
