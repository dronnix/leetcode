package main

import "fmt"

func maxDistToClosest(seats []int) int {
	start := 0 // First corner case: no neighbors on the beginning of the row:
	for ; seats[start] == 0; start++ {
	}
	max := start

	end := len(seats) - 1 // Last corner case: no neighbors on the ending of the row:
	for ; seats[end] == 0; end-- {
	}
	if len(seats)-1-end > max {
		max = len(seats) - 1 - end
	}

	prev := start // Main part: between persons:
	for i := start + 1; i <= end; i++ {
		if seats[i] == 0 {
			continue
		}
		dist := (i - prev) / 2
		if dist > max {
			max = dist
		}
		prev = i
	}

	return max
}

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
}
