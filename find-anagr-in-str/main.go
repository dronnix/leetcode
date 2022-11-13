package main

import "fmt"

// The task: https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s, p string) []int {
	matcher := NewMatcher(p)
	l := len(p) - 1
	for i := 0; i < l && i < len(s); i++ {
		matcher.Add(rune(s[i]))
	}

	res := make([]int, 0)
	for i := l; i < len(s); i++ {
		if matcher.Add(rune(s[i])) {
			res = append(res, i-l)
		}
		matcher.Remove(rune(s[i-l]))
	}
	return res
}

type Matcher struct {
	sample          map[rune]int
	current         map[rune]int
	matchCounter    int
	matchesExpected int
}

func NewMatcher(smpl string) *Matcher {
	m := &Matcher{
		sample:          make(map[rune]int, len(smpl)),
		current:         make(map[rune]int, len(smpl)),
		matchesExpected: len(smpl),
	}
	for _, c := range smpl {
		m.sample[c]++
	}
	return m
}

func (m *Matcher) Add(c rune) bool {
	if _, ok := m.sample[c]; !ok {
		return false
	}
	m.current[c]++
	if m.current[c] <= m.sample[c] {
		m.matchCounter++
	}
	return m.matchCounter == m.matchesExpected
}

func (m *Matcher) Remove(c rune) {
	if v, ok := m.current[c]; ok && v > 0 {
		m.current[c]--
		if m.current[c] < m.sample[c] {
			m.matchCounter--
		}
	}
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
