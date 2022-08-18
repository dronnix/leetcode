package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	resSize := len(matrix) * len(matrix[0])
	result := make([]int, 0, resSize)
	minI, minJ := 0, 0
	maxI, maxJ := len(matrix), len(matrix[0])

	for {
		for j := minJ; j < maxJ; j++ {
			result = append(result, matrix[minI][j])
		}
		if len(result) == resSize {
			return result
		}
		minI++
		for i := minI; i < maxI; i++ {
			result = append(result, matrix[i][maxJ-1])
		}
		if len(result) == resSize {
			return result
		}
		maxJ--
		for j := maxJ - 1; j >= minJ; j-- {
			result = append(result, matrix[maxI-1][j])
		}
		if len(result) == resSize {
			return result
		}
		maxI--
		for i := maxI - 1; i >= minI; i-- {
			result = append(result, matrix[i][minJ])
		}
		if len(result) == resSize {
			return result
		}
		minJ++
	}
}

func main() {

}
