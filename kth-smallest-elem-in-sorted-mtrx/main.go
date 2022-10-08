package main

import (
	"container/heap"
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	var h Heap = make([]Elem, 0, len(matrix))
	for i := range matrix {
		heap.Push(&h, Elem{matrix[i][0], i, 0})
	}

	for ; k > 1; k-- {
		e := heap.Pop(&h).(Elem)
		col := e.Col + 1
		if col < len(matrix[0]) {
			heap.Push(&h, Elem{matrix[e.Row][col], e.Row, col})
		}
	}
	e := heap.Pop(&h).(Elem)
	return e.Val
}

func main() {
	m := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(m, 9))
}

type Elem struct {
	Val int
	Row int
	Col int
}

type Heap []Elem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Elem))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
