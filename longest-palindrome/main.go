package main

import "fmt"

// Task: https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	var seen [140]bool
	result := 0
	notPaired := 0
	for _, c := range s {
		b := seen[c]
		if b {
			seen[c] = false
			result += 2
			notPaired--
		} else {
			seen[c] = true
			notPaired++
		}
	}
	if notPaired > 0 {
		result++
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
