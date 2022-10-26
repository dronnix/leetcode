package main

import (
	"container/heap"
	"fmt"
)

// The task: https://leetcode.com/problems/stock-price-fluctuation/

type StockPrice struct {
	last     stockQuote
	max      *StockQuoteHeap
	min      *StockQuoteHeap
	storedTS map[int]int
}

const maxCorrections = 128

func Constructor() StockPrice {
	max := StockQuoteHeap(make([]stockQuote, 0, maxCorrections))
	min := StockQuoteHeap(make([]stockQuote, 0, maxCorrections))
	stored := make(map[int]int, 2*maxCorrections)
	return StockPrice{max: &max, min: &min, storedTS: stored}

}

func (p *StockPrice) Update(timestamp int, price int) { //nolint:gocritic
	if timestamp >= p.last.ts {
		p.last = stockQuote{ts: timestamp, price: price}
	}
	heap.Push(p.max, stockQuote{ts: timestamp, price: price})
	heap.Push(p.min, stockQuote{ts: timestamp, price: -price})
	p.storedTS[timestamp] = price
}

func (p *StockPrice) Current() int {
	return p.last.price
}

func (p *StockPrice) Maximum() int {
	top := p.max.Top()
	price := p.storedTS[top.ts]
	if price > 0 && price != top.price {
		heap.Pop(p.max)
		return p.Maximum()
	}
	return top.price
}

func (p *StockPrice) Minimum() int {
	top := p.min.Top()
	price := p.storedTS[top.ts]
	if price > 0 && price != -top.price {
		heap.Pop(p.min)
		return p.Minimum()
	}
	return -top.price
}

type stockQuote struct {
	ts    int
	price int
}

type StockQuoteHeap []stockQuote

func (h StockQuoteHeap) Len() int           { return len(h) }
func (h StockQuoteHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h StockQuoteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StockQuoteHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(stockQuote))
}

func (h *StockQuoteHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h StockQuoteHeap) Top() stockQuote {
	return h[0]
}

func main() {
	sp := Constructor()
	sp.Update(1, 1)
	sp.Update(2, 2)
	sp.Update(3, 3)
	sp.Update(1, 4)
	sp.Minimum()

	fmt.Println(sp.Current(), sp.Maximum(), sp.Minimum())
}
