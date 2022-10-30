package main

import "fmt"

// The task: https://leetcode.com/problems/asteroid-collision/

func asteroidCollision(asteroids []int) []int {
	s := make([]int, 0, len(asteroids))
	for _, a := range asteroids {
		s = collide(a, s)
	}
	return s
}

func collide(a int, s []int) []int {
	// no asteroids, or left-moving only, or a is moving to the right:
	if len(s) == 0 || s[len(s)-1] < 0 || a > 0 {
		s = append(s, a)
		return s
	}

	if -a < s[len(s)-1] { // Asteroid in the stack has greater mass:
		return s
	}

	p := s[len(s)-1]
	s = s[:len(s)-1]
	if p == -a { // the same mass boss are exploded
		return s
	}
	return collide(a, s) // new asteroid wins, lets test it against others in stack!
}

func main() {
	fmt.Println(asteroidCollision([]int{-5, 2, 3, -4, 1}))
}
