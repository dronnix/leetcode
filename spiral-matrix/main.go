package main

// task: https://leetcode.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	cellsNum := len(matrix) * len(matrix[0])
	result := make([]int, 0, cellsNum)

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	steps := []int{len(matrix[0]), len(matrix) - 1, len(matrix[0]) - 1, len(matrix) - 2}
	row, col := 0, -1
	di := 0
	for steps[di] > 0 {
		for si := 0; si < steps[di]; si++ {
			row += directions[di][0]
			col += directions[di][1]
			result = append(result, matrix[row][col])
		}
		steps[di] -= 2
		di = (di + 1) % len(directions)
	}
	return result
}

func main() {

}
