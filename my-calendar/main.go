package main

import "fmt"

// The task: https://leetcode.com/problems/my-calendar-i/

type MyCalendar struct {
	slots [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{slots: make([][2]int, 0, 1000)}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, s := range c.slots {
		if overlap(s[0], start, s[1], end) {
			return false
		}
	}
	c.slots = append(c.slots, [2]int{start, end})
	return true
}

func overlap(b1, b2, e1, e2 int) bool {
	if b1 >= e2 || b2 >= e1 {
		return false
	}
	return true
}

func main() {
	c := Constructor()
	fmt.Println(c.Book(5, 10))
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(15, 17))

}
