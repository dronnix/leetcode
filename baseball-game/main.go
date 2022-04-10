package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) int {
	scores := make([]int, 0, len(ops))
	for _, o := range ops {
		l := len(scores)
		switch o {
		case "+":
			scores = append(scores, scores[l-1]+scores[l-2])
		case "D":
			scores = append(scores, scores[l-1]*2)
		case "C":
			scores = scores[:l-1]
		default:
			x, err := strconv.Atoi(o)
			if err != nil {
				panic("unexpected opt: " + o + ": " + err.Error())
			}
			scores = append(scores, x)
		}
	}

	sum := 0
	for _, s := range scores {
		sum += s
	}

	return sum
}
