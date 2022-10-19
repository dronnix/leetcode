package main

import "fmt"

// The task: https://leetcode.com/problems/minimum-genetic-mutation/

const unreachable = 65535

// TODO: Use BFS instead of DFS!

func minMutation(start, end string, bank []string) int {
	binBank := make([]uint16, 0, len(bank))
	endAtBank := false
	for i := range bank {
		binBank = append(binBank, encodeGene(bank[i]))
		if end == bank[i] {
			endAtBank = true
		}
	}
	if !endAtBank {
		return -1
	}

	m := calcMutations(encodeGene(start), encodeGene(end), binBank, 0)
	if m >= unreachable {
		return -1
	}
	return m
}

func calcMutations(start, end uint16, bank []uint16, mutationsDone int) int {
	if isOneMutation(start, end) {
		return mutationsDone + 1
	}
	min := unreachable
	for i, g := range bank {
		if isOneMutation(start, g) {
			bank[i], bank[len(bank)-1] = bank[len(bank)-1], bank[i]
			m := calcMutations(g, end, bank[:len(bank)-1], mutationsDone+1)
			if m < min {
				min = m
			}
			bank[i], bank[len(bank)-1] = bank[len(bank)-1], bank[i]
		}
	}
	return min
}

func encodeGene(g string) (r uint16) {
	for _, s := range g {
		r <<= 2
		switch s {
		case 'A':
			// leave 0
		case 'C':
			r += 1
		case 'G':
			r += 2
		case 'T':
			r += 3
		default:
			panic("unknown gene")
		}
	}
	return r
}

func isOneMutation(g1, g2 uint16) (r bool) {
	d := g1 ^ g2
	m := uint16(3)
	for i := 0; i < 8; i++ {
		if d&m > 0 {
			if r {
				return false
			}
			r = true
		}
		m <<= 2
	}
	return r
}

//"AACCGGTT"
//"AAACGGTA"
//["AACCGGTA","AACCGCTA","AAACGGTA"]

func main() {
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}
