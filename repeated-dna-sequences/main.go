package main

import "fmt"

func encodeSymbol(c rune) uint32 {
	switch c {
	case 'A':
		return 0b00
	case 'C':
		return 0b01
	case 'G':
		return 0b10
	case 'T':
		return 0b11
	}
	panic("unexpected symbol")
}

func encodeWord(s string) (result uint32) {
	for i := 0; i < 10; i++ {
		result <<= 2
		result |= encodeSymbol(rune(s[i]))
	}
	return result
}

func rollingEncodeWorld(prev uint32, c rune) uint32 {
	result := prev << 2
	result |= encodeSymbol(c)
	return result & 0x000FFFFF
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}
	m := make(map[uint32]uint8)
	r := make([]string, 0)
	w := encodeWord(s[:10])
	m[w] = 1
	for i := 10; i < len(s); i++ {
		w = rollingEncodeWorld(w, rune(s[i]))
		if m[w] > 1 {
			continue
		}
		if m[w] == 1 {
			r = append(r, s[i-9:i+1])
		}
		m[w]++
	}
	return r
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
