package main

import (
	"container/heap"
	"fmt"
)

// The task: https://leetcode.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	below MaxHeap
	above MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (f *MedianFinder) AddNum(num int) {
	if f.below.Len() == 0 || (f.below.Len() != 0 && num <= f.below.BaseHeap[0]) {
		heap.Push(&f.below, num)
		if f.below.Len()-f.above.Len() > 1 {
			heap.Push(&f.above, heap.Pop(&f.below))
		}
		return
	}
	heap.Push(&f.above, num)
	if f.above.Len()-f.below.Len() > 1 {
		heap.Push(&f.below, heap.Pop(&f.above))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	switch {
	case f.below.Len() > f.above.Len():
		return float64(f.below.BaseHeap[0])
	case f.below.Len() < f.above.Len():
		return float64(f.above.BaseHeap[0])
	default:
		return float64(f.below.BaseHeap[0]+f.above.BaseHeap[0]) / 2
	}
}

func main() {
	mf := Constructor()
	mf.AddNum(2)
	//mf.AddNum(4)
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
}

type MinHeap struct{ BaseHeap }

func (h MinHeap) Less(i, j int) bool { return h.BaseHeap[i] < h.BaseHeap[j] }

type MaxHeap struct{ BaseHeap }

func (h MaxHeap) Less(i, j int) bool { return h.BaseHeap[i] > h.BaseHeap[j] }

type BaseHeap []int

func (h BaseHeap) Len() int      { return len(h) }
func (h BaseHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *BaseHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *BaseHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
