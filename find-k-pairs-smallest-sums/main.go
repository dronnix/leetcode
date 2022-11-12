package main

import (
	"container/heap"
	"fmt"
)

// The task: https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

func kSmallestPairs(nums1, nums2 []int, k int) (res [][]int) {
	var h Heap = make([]Elem, 0, k)
	heap.Push(&h, Elem{Idx1: 0, Idx2: 0, Sum: nums1[0] + nums2[0]})
	mem := make(map[int]bool, k)
	for len(res) < k && h.Len() > 0 {
		e := heap.Pop(&h).(Elem)
		res = append(res, []int{nums1[e.Idx1], nums2[e.Idx2]})
		if e.Idx1 < len(nums1)-1 && !mem[(e.Idx1+1)*10000+e.Idx2] {
			heap.Push(&h, Elem{Idx1: e.Idx1 + 1, Idx2: e.Idx2, Sum: nums1[e.Idx1+1] + nums2[e.Idx2]})
			mem[(e.Idx1+1)*10000+e.Idx2] = true
		}
		if e.Idx2 < len(nums2)-1 && !mem[(e.Idx1)*10000+e.Idx2+1] {
			heap.Push(&h, Elem{Idx1: e.Idx1, Idx2: e.Idx2 + 1, Sum: nums1[e.Idx1] + nums2[e.Idx2+1]})
			mem[(e.Idx1)*10000+e.Idx2+1] = true
		}
	}
	return res
}

type Elem struct {
	Idx1 int
	Idx2 int
	Sum  int
}

type Heap []Elem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Sum < h[j].Sum }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Elem))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(kSmallestPairs([]int{1, 3, 7}, []int{2, 6, 8}, 7))
}
