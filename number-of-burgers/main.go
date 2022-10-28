package main

import "fmt"

// The task: https://leetcode.com/problems/number-of-burgers-with-no-waste-of-ingredients/

func numOfBurgers(tomatoSlices, cheeseSlices int) []int {
	reminder := (4*cheeseSlices - tomatoSlices) % 2
	if reminder != 0 {
		return nil
	}

	small := (4*cheeseSlices - tomatoSlices) / 2
	jumbo := cheeseSlices - small

	if small < 0 || jumbo < 0 {
		return nil
	}
	return []int{jumbo, small}
}

func main() {
	fmt.Println(numOfBurgers(16, 7))
}
