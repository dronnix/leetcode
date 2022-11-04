package main

import "fmt"

func maxLength(arr []string) int {
	ai := make([]int, len(arr))
	for i := range arr {
		a := encode(arr[i])
		if a > 0 {
			ai[i] = a
		}
	}

	res := make([]int, 0, 1<<len(ai))
	m := 0
	for _, a := range ai {
		for _, r := range res {
			if a&r == 0 {
				res = append(res, a|r)
				m = max(m, length(a|r))
			}
		}
		res = append(res, a)
		m = max(m, length(a))
	}
	return m
}

func encode(s string) (r int) {
	for _, c := range s {
		if r&(1<<(c-'a')) > 0 {
			return 0
		}
		r |= 1 << (c - 'a')
	}
	return r
}

func length(a int) (r int) {
	for i := 0; i <= 26; i++ {
		if a&(1<<i) > 0 {
			r++
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}
